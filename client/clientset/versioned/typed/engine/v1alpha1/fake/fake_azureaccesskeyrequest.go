/*
Copyright 2019 The Kube Vault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureAccessKeyRequests implements AzureAccessKeyRequestInterface
type FakeAzureAccessKeyRequests struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var azureaccesskeyrequestsResource = schema.GroupVersionResource{Group: "engine.kubevault.com", Version: "v1alpha1", Resource: "azureaccesskeyrequests"}

var azureaccesskeyrequestsKind = schema.GroupVersionKind{Group: "engine.kubevault.com", Version: "v1alpha1", Kind: "AzureAccessKeyRequest"}

// Get takes name of the azureAccessKeyRequest, and returns the corresponding azureAccessKeyRequest object, and an error if there is any.
func (c *FakeAzureAccessKeyRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureaccesskeyrequestsResource, c.ns, name), &v1alpha1.AzureAccessKeyRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAccessKeyRequest), err
}

// List takes label and field selectors, and returns the list of AzureAccessKeyRequests that match those selectors.
func (c *FakeAzureAccessKeyRequests) List(opts v1.ListOptions) (result *v1alpha1.AzureAccessKeyRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureaccesskeyrequestsResource, azureaccesskeyrequestsKind, c.ns, opts), &v1alpha1.AzureAccessKeyRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureAccessKeyRequestList{ListMeta: obj.(*v1alpha1.AzureAccessKeyRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureAccessKeyRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureAccessKeyRequests.
func (c *FakeAzureAccessKeyRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureaccesskeyrequestsResource, c.ns, opts))

}

// Create takes the representation of a azureAccessKeyRequest and creates it.  Returns the server's representation of the azureAccessKeyRequest, and an error, if there is any.
func (c *FakeAzureAccessKeyRequests) Create(azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureaccesskeyrequestsResource, c.ns, azureAccessKeyRequest), &v1alpha1.AzureAccessKeyRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAccessKeyRequest), err
}

// Update takes the representation of a azureAccessKeyRequest and updates it. Returns the server's representation of the azureAccessKeyRequest, and an error, if there is any.
func (c *FakeAzureAccessKeyRequests) Update(azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureaccesskeyrequestsResource, c.ns, azureAccessKeyRequest), &v1alpha1.AzureAccessKeyRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAccessKeyRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureAccessKeyRequests) UpdateStatus(azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest) (*v1alpha1.AzureAccessKeyRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azureaccesskeyrequestsResource, "status", c.ns, azureAccessKeyRequest), &v1alpha1.AzureAccessKeyRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAccessKeyRequest), err
}

// Delete takes name of the azureAccessKeyRequest and deletes it. Returns an error if one occurs.
func (c *FakeAzureAccessKeyRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azureaccesskeyrequestsResource, c.ns, name), &v1alpha1.AzureAccessKeyRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureAccessKeyRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureaccesskeyrequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureAccessKeyRequestList{})
	return err
}

// Patch applies the patch and returns the patched azureAccessKeyRequest.
func (c *FakeAzureAccessKeyRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureaccesskeyrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureAccessKeyRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAccessKeyRequest), err
}
