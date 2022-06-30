/*
Copyright 2022 Contributors to The HwameiStor project.

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
	"context"

	v1alpha1 "github.com/hwameistor/local-storage/pkg/apis/hwameistor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLocalVolumeGroupConverts implements LocalVolumeGroupConvertInterface
type FakeLocalVolumeGroupConverts struct {
	Fake *FakeHwameistorV1alpha1
}

var localvolumegroupconvertsResource = schema.GroupVersionResource{Group: "hwameistor.io", Version: "v1alpha1", Resource: "localvolumegroupconverts"}

var localvolumegroupconvertsKind = schema.GroupVersionKind{Group: "hwameistor.io", Version: "v1alpha1", Kind: "LocalVolumeGroupConvert"}

// Get takes name of the localVolumeGroupConvert, and returns the corresponding localVolumeGroupConvert object, and an error if there is any.
func (c *FakeLocalVolumeGroupConverts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalVolumeGroupConvert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(localvolumegroupconvertsResource, name), &v1alpha1.LocalVolumeGroupConvert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeGroupConvert), err
}

// List takes label and field selectors, and returns the list of LocalVolumeGroupConverts that match those selectors.
func (c *FakeLocalVolumeGroupConverts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalVolumeGroupConvertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(localvolumegroupconvertsResource, localvolumegroupconvertsKind, opts), &v1alpha1.LocalVolumeGroupConvertList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocalVolumeGroupConvertList{ListMeta: obj.(*v1alpha1.LocalVolumeGroupConvertList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocalVolumeGroupConvertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localVolumeGroupConverts.
func (c *FakeLocalVolumeGroupConverts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(localvolumegroupconvertsResource, opts))
}

// Create takes the representation of a localVolumeGroupConvert and creates it.  Returns the server's representation of the localVolumeGroupConvert, and an error, if there is any.
func (c *FakeLocalVolumeGroupConverts) Create(ctx context.Context, localVolumeGroupConvert *v1alpha1.LocalVolumeGroupConvert, opts v1.CreateOptions) (result *v1alpha1.LocalVolumeGroupConvert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(localvolumegroupconvertsResource, localVolumeGroupConvert), &v1alpha1.LocalVolumeGroupConvert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeGroupConvert), err
}

// Update takes the representation of a localVolumeGroupConvert and updates it. Returns the server's representation of the localVolumeGroupConvert, and an error, if there is any.
func (c *FakeLocalVolumeGroupConverts) Update(ctx context.Context, localVolumeGroupConvert *v1alpha1.LocalVolumeGroupConvert, opts v1.UpdateOptions) (result *v1alpha1.LocalVolumeGroupConvert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(localvolumegroupconvertsResource, localVolumeGroupConvert), &v1alpha1.LocalVolumeGroupConvert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeGroupConvert), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocalVolumeGroupConverts) UpdateStatus(ctx context.Context, localVolumeGroupConvert *v1alpha1.LocalVolumeGroupConvert, opts v1.UpdateOptions) (*v1alpha1.LocalVolumeGroupConvert, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(localvolumegroupconvertsResource, "status", localVolumeGroupConvert), &v1alpha1.LocalVolumeGroupConvert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeGroupConvert), err
}

// Delete takes name of the localVolumeGroupConvert and deletes it. Returns an error if one occurs.
func (c *FakeLocalVolumeGroupConverts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(localvolumegroupconvertsResource, name), &v1alpha1.LocalVolumeGroupConvert{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalVolumeGroupConverts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(localvolumegroupconvertsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocalVolumeGroupConvertList{})
	return err
}

// Patch applies the patch and returns the patched localVolumeGroupConvert.
func (c *FakeLocalVolumeGroupConverts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalVolumeGroupConvert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(localvolumegroupconvertsResource, name, pt, data, subresources...), &v1alpha1.LocalVolumeGroupConvert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeGroupConvert), err
}
