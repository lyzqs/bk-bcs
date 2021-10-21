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

// Code generated by apiregister-gen. DO NOT EDIT.

package apis

import (
	"github.com/Tencent/bk-bcs/bcs-k8s/bcs-federated-apiserver/pkg/apis/aggregation"
	_ "github.com/Tencent/bk-bcs/bcs-k8s/bcs-federated-apiserver/pkg/apis/aggregation/install" // Install the aggregation group
	aggregationv1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-federated-apiserver/pkg/apis/aggregation/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		aggregationv1alpha1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetAggregationAPIBuilder(),
	}
}

var aggregationApiGroup = builders.NewApiGroupBuilder(
	"aggregation.federated.bkbcs.tencent.com",
	"github.com/Tencent/bk-bcs/bcs-k8s/bcs-federated-apiserver/pkg/apis/aggregation").
	WithUnVersionedApi(aggregation.ApiVersion).
	WithVersionedApis(
		aggregationv1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetAggregationAPIBuilder() *builders.APIGroupBuilder {
	return aggregationApiGroup
}
