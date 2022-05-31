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

// FakeLocalVolumeExpands implements LocalVolumeExpandInterface
type FakeLocalVolumeExpands struct {
	Fake *FakeHwameistorV1alpha1
}

var localvolumeexpandsResource = schema.GroupVersionResource{Group: "hwameistor.io", Version: "v1alpha1", Resource: "localvolumeexpands"}

var localvolumeexpandsKind = schema.GroupVersionKind{Group: "hwameistor.io", Version: "v1alpha1", Kind: "LocalVolumeExpand"}

// Get takes name of the localVolumeExpand, and returns the corresponding localVolumeExpand object, and an error if there is any.
func (c *FakeLocalVolumeExpands) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalVolumeExpand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(localvolumeexpandsResource, name), &v1alpha1.LocalVolumeExpand{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeExpand), err
}

// List takes label and field selectors, and returns the list of LocalVolumeExpands that match those selectors.
func (c *FakeLocalVolumeExpands) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalVolumeExpandList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(localvolumeexpandsResource, localvolumeexpandsKind, opts), &v1alpha1.LocalVolumeExpandList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocalVolumeExpandList{ListMeta: obj.(*v1alpha1.LocalVolumeExpandList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocalVolumeExpandList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localVolumeExpands.
func (c *FakeLocalVolumeExpands) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(localvolumeexpandsResource, opts))
}

// Create takes the representation of a localVolumeExpand and creates it.  Returns the server's representation of the localVolumeExpand, and an error, if there is any.
func (c *FakeLocalVolumeExpands) Create(ctx context.Context, localVolumeExpand *v1alpha1.LocalVolumeExpand, opts v1.CreateOptions) (result *v1alpha1.LocalVolumeExpand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(localvolumeexpandsResource, localVolumeExpand), &v1alpha1.LocalVolumeExpand{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeExpand), err
}

// Update takes the representation of a localVolumeExpand and updates it. Returns the server's representation of the localVolumeExpand, and an error, if there is any.
func (c *FakeLocalVolumeExpands) Update(ctx context.Context, localVolumeExpand *v1alpha1.LocalVolumeExpand, opts v1.UpdateOptions) (result *v1alpha1.LocalVolumeExpand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(localvolumeexpandsResource, localVolumeExpand), &v1alpha1.LocalVolumeExpand{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeExpand), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocalVolumeExpands) UpdateStatus(ctx context.Context, localVolumeExpand *v1alpha1.LocalVolumeExpand, opts v1.UpdateOptions) (*v1alpha1.LocalVolumeExpand, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(localvolumeexpandsResource, "status", localVolumeExpand), &v1alpha1.LocalVolumeExpand{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeExpand), err
}

// Delete takes name of the localVolumeExpand and deletes it. Returns an error if one occurs.
func (c *FakeLocalVolumeExpands) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(localvolumeexpandsResource, name), &v1alpha1.LocalVolumeExpand{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalVolumeExpands) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(localvolumeexpandsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocalVolumeExpandList{})
	return err
}

// Patch applies the patch and returns the patched localVolumeExpand.
func (c *FakeLocalVolumeExpands) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalVolumeExpand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(localvolumeexpandsResource, name, pt, data, subresources...), &v1alpha1.LocalVolumeExpand{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalVolumeExpand), err
}
