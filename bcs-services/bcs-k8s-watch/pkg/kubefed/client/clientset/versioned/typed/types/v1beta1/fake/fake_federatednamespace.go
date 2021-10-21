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
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedNamespaces implements FederatedNamespaceInterface
type FakeFederatedNamespaces struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatednamespacesResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatednamespaces"}

var federatednamespacesKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedNamespace"}

// Get takes name of the federatedNamespace, and returns the corresponding federatedNamespace object, and an error if there is any.
func (c *FakeFederatedNamespaces) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatednamespacesResource, c.ns, name), &v1beta1.FederatedNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedNamespace), err
}

// List takes label and field selectors, and returns the list of FederatedNamespaces that match those selectors.
func (c *FakeFederatedNamespaces) List(opts v1.ListOptions) (result *v1beta1.FederatedNamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatednamespacesResource, federatednamespacesKind, c.ns, opts), &v1beta1.FederatedNamespaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedNamespaceList{}
	for _, item := range obj.(*v1beta1.FederatedNamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedNamespaces.
func (c *FakeFederatedNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatednamespacesResource, c.ns, opts))

}

// Create takes the representation of a federatedNamespace and creates it.  Returns the server's representation of the federatedNamespace, and an error, if there is any.
func (c *FakeFederatedNamespaces) Create(federatedNamespace *v1beta1.FederatedNamespace) (result *v1beta1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatednamespacesResource, c.ns, federatedNamespace), &v1beta1.FederatedNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedNamespace), err
}

// Update takes the representation of a federatedNamespace and updates it. Returns the server's representation of the federatedNamespace, and an error, if there is any.
func (c *FakeFederatedNamespaces) Update(federatedNamespace *v1beta1.FederatedNamespace) (result *v1beta1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatednamespacesResource, c.ns, federatedNamespace), &v1beta1.FederatedNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedNamespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedNamespaces) UpdateStatus(federatedNamespace *v1beta1.FederatedNamespace) (*v1beta1.FederatedNamespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatednamespacesResource, "status", c.ns, federatedNamespace), &v1beta1.FederatedNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedNamespace), err
}

// Delete takes name of the federatedNamespace and deletes it. Returns an error if one occurs.
func (c *FakeFederatedNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatednamespacesResource, c.ns, name), &v1beta1.FederatedNamespace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatednamespacesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedNamespaceList{})
	return err
}

// Patch applies the patch and returns the patched federatedNamespace.
func (c *FakeFederatedNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatednamespacesResource, c.ns, name, data, subresources...), &v1beta1.FederatedNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedNamespace), err
}
