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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/core/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubeFedConfigLister helps list KubeFedConfigs.
type KubeFedConfigLister interface {
	// List lists all KubeFedConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.KubeFedConfig, err error)
	// KubeFedConfigs returns an object that can list and get KubeFedConfigs.
	KubeFedConfigs(namespace string) KubeFedConfigNamespaceLister
	KubeFedConfigListerExpansion
}

// kubeFedConfigLister implements the KubeFedConfigLister interface.
type kubeFedConfigLister struct {
	indexer cache.Indexer
}

// NewKubeFedConfigLister returns a new KubeFedConfigLister.
func NewKubeFedConfigLister(indexer cache.Indexer) KubeFedConfigLister {
	return &kubeFedConfigLister{indexer: indexer}
}

// List lists all KubeFedConfigs in the indexer.
func (s *kubeFedConfigLister) List(selector labels.Selector) (ret []*v1beta1.KubeFedConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.KubeFedConfig))
	})
	return ret, err
}

// KubeFedConfigs returns an object that can list and get KubeFedConfigs.
func (s *kubeFedConfigLister) KubeFedConfigs(namespace string) KubeFedConfigNamespaceLister {
	return kubeFedConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubeFedConfigNamespaceLister helps list and get KubeFedConfigs.
type KubeFedConfigNamespaceLister interface {
	// List lists all KubeFedConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.KubeFedConfig, err error)
	// Get retrieves the KubeFedConfig from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.KubeFedConfig, error)
	KubeFedConfigNamespaceListerExpansion
}

// kubeFedConfigNamespaceLister implements the KubeFedConfigNamespaceLister
// interface.
type kubeFedConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KubeFedConfigs in the indexer for a given namespace.
func (s kubeFedConfigNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.KubeFedConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.KubeFedConfig))
	})
	return ret, err
}

// Get retrieves the KubeFedConfig from the indexer for a given namespace and name.
func (s kubeFedConfigNamespaceLister) Get(name string) (*v1beta1.KubeFedConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("kubefedconfig"), name)
	}
	return obj.(*v1beta1.KubeFedConfig), nil
}
