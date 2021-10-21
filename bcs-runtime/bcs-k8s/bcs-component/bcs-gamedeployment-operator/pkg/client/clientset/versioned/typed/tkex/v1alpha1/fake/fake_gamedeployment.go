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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-gamedeployment-operator/pkg/apis/tkex/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGameDeployments implements GameDeploymentInterface
type FakeGameDeployments struct {
	Fake *FakeTkexV1alpha1
	ns   string
}

var gamedeploymentsResource = schema.GroupVersionResource{Group: "tkex", Version: "v1alpha1", Resource: "gamedeployments"}

var gamedeploymentsKind = schema.GroupVersionKind{Group: "tkex", Version: "v1alpha1", Kind: "GameDeployment"}

// Get takes name of the gameDeployment, and returns the corresponding gameDeployment object, and an error if there is any.
func (c *FakeGameDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.GameDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gamedeploymentsResource, c.ns, name), &v1alpha1.GameDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameDeployment), err
}

// List takes label and field selectors, and returns the list of GameDeployments that match those selectors.
func (c *FakeGameDeployments) List(opts v1.ListOptions) (result *v1alpha1.GameDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gamedeploymentsResource, gamedeploymentsKind, c.ns, opts), &v1alpha1.GameDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GameDeploymentList{ListMeta: obj.(*v1alpha1.GameDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.GameDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gameDeployments.
func (c *FakeGameDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gamedeploymentsResource, c.ns, opts))

}

// Create takes the representation of a gameDeployment and creates it.  Returns the server's representation of the gameDeployment, and an error, if there is any.
func (c *FakeGameDeployments) Create(gameDeployment *v1alpha1.GameDeployment) (result *v1alpha1.GameDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gamedeploymentsResource, c.ns, gameDeployment), &v1alpha1.GameDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameDeployment), err
}

// Update takes the representation of a gameDeployment and updates it. Returns the server's representation of the gameDeployment, and an error, if there is any.
func (c *FakeGameDeployments) Update(gameDeployment *v1alpha1.GameDeployment) (result *v1alpha1.GameDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gamedeploymentsResource, c.ns, gameDeployment), &v1alpha1.GameDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGameDeployments) UpdateStatus(gameDeployment *v1alpha1.GameDeployment) (*v1alpha1.GameDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gamedeploymentsResource, "status", c.ns, gameDeployment), &v1alpha1.GameDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameDeployment), err
}

// Delete takes name of the gameDeployment and deletes it. Returns an error if one occurs.
func (c *FakeGameDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gamedeploymentsResource, c.ns, name), &v1alpha1.GameDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGameDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gamedeploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GameDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched gameDeployment.
func (c *FakeGameDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gamedeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GameDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameDeployment), err
}
