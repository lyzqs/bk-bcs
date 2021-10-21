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
	"context"

	cloudv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/cloud/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudIPQuotas implements CloudIPQuotaInterface
type FakeCloudIPQuotas struct {
	Fake *FakeCloudV1
	ns   string
}

var cloudipquotasResource = schema.GroupVersionResource{Group: "cloud", Version: "v1", Resource: "cloudipquotas"}

var cloudipquotasKind = schema.GroupVersionKind{Group: "cloud", Version: "v1", Kind: "CloudIPQuota"}

// Get takes name of the cloudIPQuota, and returns the corresponding cloudIPQuota object, and an error if there is any.
func (c *FakeCloudIPQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *cloudv1.CloudIPQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudipquotasResource, c.ns, name), &cloudv1.CloudIPQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIPQuota), err
}

// List takes label and field selectors, and returns the list of CloudIPQuotas that match those selectors.
func (c *FakeCloudIPQuotas) List(ctx context.Context, opts v1.ListOptions) (result *cloudv1.CloudIPQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudipquotasResource, cloudipquotasKind, c.ns, opts), &cloudv1.CloudIPQuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudv1.CloudIPQuotaList{ListMeta: obj.(*cloudv1.CloudIPQuotaList).ListMeta}
	for _, item := range obj.(*cloudv1.CloudIPQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudIPQuotas.
func (c *FakeCloudIPQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudipquotasResource, c.ns, opts))

}

// Create takes the representation of a cloudIPQuota and creates it.  Returns the server's representation of the cloudIPQuota, and an error, if there is any.
func (c *FakeCloudIPQuotas) Create(ctx context.Context, cloudIPQuota *cloudv1.CloudIPQuota, opts v1.CreateOptions) (result *cloudv1.CloudIPQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudipquotasResource, c.ns, cloudIPQuota), &cloudv1.CloudIPQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIPQuota), err
}

// Update takes the representation of a cloudIPQuota and updates it. Returns the server's representation of the cloudIPQuota, and an error, if there is any.
func (c *FakeCloudIPQuotas) Update(ctx context.Context, cloudIPQuota *cloudv1.CloudIPQuota, opts v1.UpdateOptions) (result *cloudv1.CloudIPQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudipquotasResource, c.ns, cloudIPQuota), &cloudv1.CloudIPQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIPQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudIPQuotas) UpdateStatus(ctx context.Context, cloudIPQuota *cloudv1.CloudIPQuota, opts v1.UpdateOptions) (*cloudv1.CloudIPQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudipquotasResource, "status", c.ns, cloudIPQuota), &cloudv1.CloudIPQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIPQuota), err
}

// Delete takes name of the cloudIPQuota and deletes it. Returns an error if one occurs.
func (c *FakeCloudIPQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudipquotasResource, c.ns, name), &cloudv1.CloudIPQuota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudIPQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudipquotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &cloudv1.CloudIPQuotaList{})
	return err
}

// Patch applies the patch and returns the patched cloudIPQuota.
func (c *FakeCloudIPQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cloudv1.CloudIPQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudipquotasResource, c.ns, name, pt, data, subresources...), &cloudv1.CloudIPQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIPQuota), err
}
