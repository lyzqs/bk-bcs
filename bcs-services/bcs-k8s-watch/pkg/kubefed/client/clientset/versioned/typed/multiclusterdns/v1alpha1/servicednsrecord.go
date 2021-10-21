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

package v1alpha1

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/multiclusterdns/v1alpha1"
	scheme "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceDNSRecordsGetter has a method to return a ServiceDNSRecordInterface.
// A group's client should implement this interface.
type ServiceDNSRecordsGetter interface {
	ServiceDNSRecords(namespace string) ServiceDNSRecordInterface
}

// ServiceDNSRecordInterface has methods to work with ServiceDNSRecord resources.
type ServiceDNSRecordInterface interface {
	Create(*v1alpha1.ServiceDNSRecord) (*v1alpha1.ServiceDNSRecord, error)
	Update(*v1alpha1.ServiceDNSRecord) (*v1alpha1.ServiceDNSRecord, error)
	UpdateStatus(*v1alpha1.ServiceDNSRecord) (*v1alpha1.ServiceDNSRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceDNSRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceDNSRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDNSRecord, err error)
	ServiceDNSRecordExpansion
}

// serviceDNSRecords implements ServiceDNSRecordInterface
type serviceDNSRecords struct {
	client rest.Interface
	ns     string
}

// newServiceDNSRecords returns a ServiceDNSRecords
func newServiceDNSRecords(c *MulticlusterdnsV1alpha1Client, namespace string) *serviceDNSRecords {
	return &serviceDNSRecords{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceDNSRecord, and returns the corresponding serviceDNSRecord object, and an error if there is any.
func (c *serviceDNSRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceDNSRecord, err error) {
	result = &v1alpha1.ServiceDNSRecord{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicednsrecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceDNSRecords that match those selectors.
func (c *serviceDNSRecords) List(opts v1.ListOptions) (result *v1alpha1.ServiceDNSRecordList, err error) {
	result = &v1alpha1.ServiceDNSRecordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicednsrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceDNSRecords.
func (c *serviceDNSRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicednsrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a serviceDNSRecord and creates it.  Returns the server's representation of the serviceDNSRecord, and an error, if there is any.
func (c *serviceDNSRecords) Create(serviceDNSRecord *v1alpha1.ServiceDNSRecord) (result *v1alpha1.ServiceDNSRecord, err error) {
	result = &v1alpha1.ServiceDNSRecord{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicednsrecords").
		Body(serviceDNSRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceDNSRecord and updates it. Returns the server's representation of the serviceDNSRecord, and an error, if there is any.
func (c *serviceDNSRecords) Update(serviceDNSRecord *v1alpha1.ServiceDNSRecord) (result *v1alpha1.ServiceDNSRecord, err error) {
	result = &v1alpha1.ServiceDNSRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicednsrecords").
		Name(serviceDNSRecord.Name).
		Body(serviceDNSRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *serviceDNSRecords) UpdateStatus(serviceDNSRecord *v1alpha1.ServiceDNSRecord) (result *v1alpha1.ServiceDNSRecord, err error) {
	result = &v1alpha1.ServiceDNSRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicednsrecords").
		Name(serviceDNSRecord.Name).
		SubResource("status").
		Body(serviceDNSRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceDNSRecord and deletes it. Returns an error if one occurs.
func (c *serviceDNSRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicednsrecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceDNSRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicednsrecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceDNSRecord.
func (c *serviceDNSRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDNSRecord, err error) {
	result = &v1alpha1.ServiceDNSRecord{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicednsrecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
