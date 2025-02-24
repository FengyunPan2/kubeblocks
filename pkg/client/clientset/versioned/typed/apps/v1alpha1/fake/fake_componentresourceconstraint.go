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

// FakeComponentResourceConstraints implements ComponentResourceConstraintInterface
type FakeComponentResourceConstraints struct {
	Fake *FakeAppsV1alpha1
}

var componentresourceconstraintsResource = schema.GroupVersionResource{Group: "apps.kubeblocks.io", Version: "v1alpha1", Resource: "componentresourceconstraints"}

var componentresourceconstraintsKind = schema.GroupVersionKind{Group: "apps.kubeblocks.io", Version: "v1alpha1", Kind: "ComponentResourceConstraint"}

// Get takes name of the componentResourceConstraint, and returns the corresponding componentResourceConstraint object, and an error if there is any.
func (c *FakeComponentResourceConstraints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComponentResourceConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(componentresourceconstraintsResource, name), &v1alpha1.ComponentResourceConstraint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentResourceConstraint), err
}

// List takes label and field selectors, and returns the list of ComponentResourceConstraints that match those selectors.
func (c *FakeComponentResourceConstraints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComponentResourceConstraintList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(componentresourceconstraintsResource, componentresourceconstraintsKind, opts), &v1alpha1.ComponentResourceConstraintList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComponentResourceConstraintList{ListMeta: obj.(*v1alpha1.ComponentResourceConstraintList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComponentResourceConstraintList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested componentResourceConstraints.
func (c *FakeComponentResourceConstraints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(componentresourceconstraintsResource, opts))
}

// Create takes the representation of a componentResourceConstraint and creates it.  Returns the server's representation of the componentResourceConstraint, and an error, if there is any.
func (c *FakeComponentResourceConstraints) Create(ctx context.Context, componentResourceConstraint *v1alpha1.ComponentResourceConstraint, opts v1.CreateOptions) (result *v1alpha1.ComponentResourceConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(componentresourceconstraintsResource, componentResourceConstraint), &v1alpha1.ComponentResourceConstraint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentResourceConstraint), err
}

// Update takes the representation of a componentResourceConstraint and updates it. Returns the server's representation of the componentResourceConstraint, and an error, if there is any.
func (c *FakeComponentResourceConstraints) Update(ctx context.Context, componentResourceConstraint *v1alpha1.ComponentResourceConstraint, opts v1.UpdateOptions) (result *v1alpha1.ComponentResourceConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(componentresourceconstraintsResource, componentResourceConstraint), &v1alpha1.ComponentResourceConstraint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentResourceConstraint), err
}

// Delete takes name of the componentResourceConstraint and deletes it. Returns an error if one occurs.
func (c *FakeComponentResourceConstraints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(componentresourceconstraintsResource, name, opts), &v1alpha1.ComponentResourceConstraint{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComponentResourceConstraints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(componentresourceconstraintsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComponentResourceConstraintList{})
	return err
}

// Patch applies the patch and returns the patched componentResourceConstraint.
func (c *FakeComponentResourceConstraints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComponentResourceConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(componentresourceconstraintsResource, name, pt, data, subresources...), &v1alpha1.ComponentResourceConstraint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComponentResourceConstraint), err
}
