/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

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

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterVersions implements ClusterVersionInterface
type FakeClusterVersions struct {
	Fake *FakeAppsV1alpha1
}

var clusterversionsResource = schema.GroupVersionResource{Group: "apps.kubeblocks.io", Version: "v1alpha1", Resource: "clusterversions"}

var clusterversionsKind = schema.GroupVersionKind{Group: "apps.kubeblocks.io", Version: "v1alpha1", Kind: "ClusterVersion"}

// Get takes name of the clusterVersion, and returns the corresponding clusterVersion object, and an error if there is any.
func (c *FakeClusterVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterversionsResource, name), &v1alpha1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// List takes label and field selectors, and returns the list of ClusterVersions that match those selectors.
func (c *FakeClusterVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterversionsResource, clusterversionsKind, opts), &v1alpha1.ClusterVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterVersionList{ListMeta: obj.(*v1alpha1.ClusterVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterVersions.
func (c *FakeClusterVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterversionsResource, opts))
}

// Create takes the representation of a clusterVersion and creates it.  Returns the server's representation of the clusterVersion, and an error, if there is any.
func (c *FakeClusterVersions) Create(ctx context.Context, clusterVersion *v1alpha1.ClusterVersion, opts v1.CreateOptions) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterversionsResource, clusterVersion), &v1alpha1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// Update takes the representation of a clusterVersion and updates it. Returns the server's representation of the clusterVersion, and an error, if there is any.
func (c *FakeClusterVersions) Update(ctx context.Context, clusterVersion *v1alpha1.ClusterVersion, opts v1.UpdateOptions) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterversionsResource, clusterVersion), &v1alpha1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterVersions) UpdateStatus(ctx context.Context, clusterVersion *v1alpha1.ClusterVersion, opts v1.UpdateOptions) (*v1alpha1.ClusterVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterversionsResource, "status", clusterVersion), &v1alpha1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}

// Delete takes name of the clusterVersion and deletes it. Returns an error if one occurs.
func (c *FakeClusterVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterversionsResource, name, opts), &v1alpha1.ClusterVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterversionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterVersionList{})
	return err
}

// Patch applies the patch and returns the patched clusterVersion.
func (c *FakeClusterVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterversionsResource, name, pt, data, subresources...), &v1alpha1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVersion), err
}
