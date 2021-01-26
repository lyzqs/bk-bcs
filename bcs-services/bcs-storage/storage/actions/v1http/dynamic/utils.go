/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package dynamic

import (
	"time"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/v2/broker"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/codec"
	"github.com/Tencent/bk-bcs/bcs-common/common/types"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/msgqueue"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/odm/operator"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/actions/lib"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/apiserver"
)

const (
	urlPrefixK8S    = "/k8s"
	urlPrefixMesos  = "/mesos"
	clusterIDTag    = "clusterId"
	namespaceTag    = "namespace"
	resourceTypeTag = "resourceType"
	resourceNameTag = "resourceName"

	tableTag      = resourceTypeTag
	dataTag       = "data"
	extraTag      = "extra"
	fieldTag      = "field"
	offsetTag     = "offset"
	limitTag      = "limit"
	updateTimeTag = "updateTime"
	createTimeTag = "createTime"

	applicationTypeName = "application"
	processTypeName     = "process"
	kindTag             = "data.kind"
)

var needTimeFormatList = []string{updateTimeTag, createTimeTag}
var nsFeatTags = []string{clusterIDTag, namespaceTag, resourceTypeTag, resourceNameTag}
var csFeatTags = []string{clusterIDTag, resourceTypeTag, resourceNameTag}
var nsListFeatTags = []string{clusterIDTag, namespaceTag, resourceTypeTag}
var csListFeatTags = []string{clusterIDTag, resourceTypeTag}
var indexKeys = []string{resourceNameTag, namespaceTag}

// Use Mongodb for storage.
const dbConfig = "mongodb/dynamic"

func getSelector(req *restful.Request) []string {
	return lib.GetQueryParamStringArray(req, fieldTag, ",")
}

func getTable(req *restful.Request) string {
	table := req.PathParameter(tableTag)
	// for mesos
	if table == processTypeName {
		table = applicationTypeName
	}
	return table
}

func getExtra(req *restful.Request) operator.M {
	raw := req.QueryParameter(extraTag)
	if raw == "" {
		return nil
	}
	extra := make(operator.M)
	lib.NewExtra(raw).Unmarshal(&extra)
	return extra
}

func getFeatures(req *restful.Request, resourceFeatList []string) operator.M {
	features := make(operator.M)
	for _, key := range resourceFeatList {
		features[key] = req.PathParameter(key)
	}
	return features
}

func getCondition(req *restful.Request, resourceFeatList []string) *operator.Condition {
	features := getFeatures(req, resourceFeatList)
	extras := getExtra(req)
	features.Merge(extras)
	featuresExcept := make(operator.M)
	for key := range features {
		// For historical reasons, mesos process is stored with application in one table(same clusters).
		// And process's construction is almost the same with application, except with field 'data.kind'.
		// If 'data.kind'='process', then this object is a process stored in application-table,
		// If 'data.kind'='application' or '', then this object is an application stored in application-table.
		//
		// For this case, we should:
		// 1. Change the key 'resourceType' from 'process' to 'application' when the caller ask for 'process'.
		// 2. Besides, getFeat() should add an extra condition that
		//    mentions the 'data.kind' to distinguish 'process' and 'application'.
		// 3. Make sure the table is application-table whether the type is 'application' or 'process'. (with getTable())
		if key == resourceTypeTag {
			switch features[key] {
			case applicationTypeName:
				featuresExcept[kindTag] = processTypeName
			case processTypeName:
				features[key] = applicationTypeName
				features[kindTag] = processTypeName
			}
		}
	}
	condition := operator.NewLeafCondition(operator.Eq, features)
	if len(featuresExcept) == 0 {
		notCondition := operator.NewLeafCondition(operator.Ne, featuresExcept)
		condition = operator.NewBranchCondition(operator.And,
			condition, notCondition)
	}
	return condition
}

func getNamespaceResources(req *restful.Request) ([]operator.M, error) {
	return getResources(req, nsFeatTags)
}

func getClusterResources(req *restful.Request) ([]operator.M, error) {
	return getResources(req, csFeatTags)
}

func listNamespaceResources(req *restful.Request) ([]operator.M, error) {
	return getResources(req, nsListFeatTags)
}

func listClusterResources(req *restful.Request) ([]operator.M, error) {
	return getResources(req, csListFeatTags)
}

func getResources(req *restful.Request, resourceFeatList []string) ([]operator.M, error) {
	condition := getCondition(req, resourceFeatList)
	offset, err := lib.GetQueryParamInt64(req, offsetTag, 0)
	if err != nil {
		return nil, err
	}
	limit, err := lib.GetQueryParamInt64(req, limitTag, 0)
	if err != nil {
		return nil, err
	}
	getOption := &lib.StoreGetOption{
		Fields: getSelector(req),
		Cond:   condition,
		Offset: offset,
		Limit:  limit,
	}
	store := lib.NewStore(
		apiserver.GetAPIResource().GetDBClient(dbConfig),
		apiserver.GetAPIResource().GetEventBus(dbConfig))
	mList, err := store.Get(req.Request.Context(), getTable(req), getOption)
	if err != nil {
		return nil, err
	}
	lib.FormatTime(mList, needTimeFormatList)
	return mList, err
}

func getReqData(req *restful.Request, features operator.M) (operator.M, error) {
	var tmp types.BcsStorageDynamicIf
	if err := codec.DecJsonReader(req.Request.Body, &tmp); err != nil {
		return nil, err
	}
	data := lib.CopyMap(features)
	data[dataTag] = tmp.Data
	return data, nil
}

func putNamespaceResources(req *restful.Request) error {
	data, err := putResources(req, nsFeatTags)
	if err != nil {
		return err
	}

	err = publishDynamicResourceToQueue(data, nsFeatTags, msgqueue.EventTypeUpdate)
	if err != nil {
		blog.Errorf("func[%s] call publishDynamicResourceToQueue failed: err[%v]", "putNamespaceResources", err)
	}

	return nil
}

func putClusterResources(req *restful.Request) error {
	data, err := putResources(req, csFeatTags)
	if err != nil {
		return err
	}

	err = publishDynamicResourceToQueue(data, csFeatTags, msgqueue.EventTypeUpdate)
	if err != nil {
		blog.Errorf("func[%s] call publishDynamicResourceToQueue failed: err[%v]", "putClusterResources", err)
	}

	return nil
}

func putResources(req *restful.Request, resourceFeatList []string) (operator.M, error) {
	features := getFeatures(req, resourceFeatList)
	extras := getExtra(req)
	features.Merge(extras)
	data, err := getReqData(req, features)
	if err != nil {
		return nil, err
	}
	putOption := &lib.StorePutOption{
		UniqueKey:     resourceFeatList,
		Cond:          operator.NewLeafCondition(operator.Eq, features),
		CreateTimeKey: createTimeTag,
		UpdateTimeKey: updateTimeTag,
	}
	store := lib.NewStore(
		apiserver.GetAPIResource().GetDBClient(dbConfig),
		apiserver.GetAPIResource().GetEventBus(dbConfig))

	err = store.Put(req.Request.Context(), getTable(req), data, putOption)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func deleteNamespaceResources(req *restful.Request) error {
	mList, err := deleteResources(req, nsFeatTags)
	if err != nil {
		return err
	}

	go func(mList []operator.M, featTags []string) {
		for _, data := range mList {
			err := publishDynamicResourceToQueue(data, featTags, msgqueue.EventTypeDelete)
			if err != nil {
				blog.Errorf("func[%s] call publishDynamicResourceToQueue failed: err[%v]", "deleteNamespaceResources", err)
			}
		}
	}(mList, nsFeatTags)

	return nil
}

func deleteClusterResources(req *restful.Request) error {
	mList, err := deleteResources(req, csFeatTags)
	if err != nil {
		return err
	}

	go func(mList []operator.M, featTags []string) {
		for _, data := range mList {
			err := publishDynamicResourceToQueue(data, featTags, msgqueue.EventTypeDelete)
			if err != nil {
				blog.Errorf("func[%s] call publishDynamicResourceToQueue failed: err[%v]", "deleteClusterResources", err)
			}
		}
	}(mList, csFeatTags)

	return nil
}

func deleteResources(req *restful.Request, resourceFeatList []string) ([]operator.M, error) {
	condition := getCondition(req, resourceFeatList)

	getOption := &lib.StoreGetOption{
		Cond:           condition,
		IsAllDocuments: true,
	}

	rmOption := &lib.StoreRemoveOption{
		Cond:           condition,
		IgnoreNotFound: false,
	}
	store := lib.NewStore(
		apiserver.GetAPIResource().GetDBClient(dbConfig),
		apiserver.GetAPIResource().GetEventBus(dbConfig))

	mList, err := store.Get(req.Request.Context(), getTable(req), getOption)
	if err != nil {
		return nil, err
	}
	lib.FormatTime(mList, needTimeFormatList)

	err = store.Remove(req.Request.Context(), getTable(req), rmOption)
	if err != nil {
		return nil, err
	}

	return mList, nil
}

func getTimeCondition(req *restful.Request) *operator.Condition {
	var data types.BcsStorageDynamicBatchDeleteIf
	if err := codec.DecJsonReader(req.Request.Body, &data); err != nil {
		return operator.EmptyCondition
	}

	condList := make([]*operator.Condition, 0)
	if data.UpdateTimeBegin > 0 {
		condList = append(condList, operator.NewLeafCondition(operator.Gt, operator.M{
			updateTimeTag: time.Unix(data.UpdateTimeBegin, 0)}))
	}
	if data.UpdateTimeEnd > 0 {
		condList = append(condList, operator.NewLeafCondition(operator.Lt, operator.M{
			updateTimeTag: time.Unix(data.UpdateTimeEnd, 0)}))
	}
	if len(condList) == 0 {
		return operator.EmptyCondition
	}
	return operator.NewBranchCondition(operator.And, condList...)
}

func deleteBatchNamespaceResource(req *restful.Request) error {
	mList, err := deleteBatchResources(req, nsListFeatTags)
	if err != nil {
		return err
	}

	go func(mList []operator.M, featTags []string) {
		for _, data := range mList {
			err := publishDynamicResourceToQueue(data, featTags, msgqueue.EventTypeDelete)
			if err != nil {
				blog.Errorf("func[%s] call publishDynamicResourceToQueue failed: err[%v]", "deleteBatchNamespaceResource", err)
			}
		}
	}(mList, nsListFeatTags)

	return nil
}

func deleteClusterNamespaceResource(req *restful.Request) error {
	mList, err := deleteBatchResources(req, csListFeatTags)
	if err != nil {
		return err
	}

	go func(mList []operator.M, featTags []string) {
		for _, data := range mList {
			err := publishDynamicResourceToQueue(data, featTags, msgqueue.EventTypeDelete)
			if err != nil {
				blog.Errorf("func[%s] call publishDynamicResourceToQueue failed: err[%v]", "deleteClusterNamespaceResource", err)
			}
		}
	}(mList, csListFeatTags)

	return nil
}

func deleteBatchResources(req *restful.Request, resourceFeatList []string) ([]operator.M, error) {
	featCondition := getCondition(req, resourceFeatList)
	timeCondition := getTimeCondition(req)
	condition := operator.NewBranchCondition(operator.And, featCondition, timeCondition)

	getOption := &lib.StoreGetOption{
		Cond:           condition,
		IsAllDocuments: true,
	}
	rmOption := &lib.StoreRemoveOption{
		Cond:           condition,
		IgnoreNotFound: false,
	}
	store := lib.NewStore(
		apiserver.GetAPIResource().GetDBClient(dbConfig),
		apiserver.GetAPIResource().GetEventBus(dbConfig))

	mList, err := store.Get(req.Request.Context(), getTable(req), getOption)
	if err != nil {
		return nil, err
	}
	lib.FormatTime(mList, needTimeFormatList)

	err = store.Remove(req.Request.Context(), getTable(req), rmOption)
	if err != nil {
		return nil, err
	}

	return mList, nil
}

func urlPathK8S(oldURL string) string {
	return urlPrefixK8S + oldURL
}

func urlPathMesos(oldURL string) string {
	return urlPrefixMesos + oldURL
}

func publishDynamicResourceToQueue(data operator.M, featTags []string, event msgqueue.EventKind) error {
	var (
		err     error
		message = &broker.Message{
			Header: map[string]string{},
		}
	)

	startTime := time.Now()
	defer func() {
		if queueName, ok := message.Header[resourceTypeTag]; ok {
			lib.ReportQueuePushMetrics(queueName, err, startTime)
		}
	}()

	for _, feat := range featTags {
		if v, ok := data[feat].(string); ok {
			message.Header[feat] = v
		}
	}
	message.Header[string(msgqueue.EventType)] = string(event)

	if v, ok := data[dataTag]; ok {
		codec.EncJson(v, &message.Body)
	} else {
		blog.Infof("object[%v] not exist data", data[dataTag])
		return nil
	}

	err = apiserver.GetAPIResource().GetMsgQueue().Publish(message)
	if err != nil {
		return err
	}

	return nil
}