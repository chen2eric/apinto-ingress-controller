/*
Copyright The Kubernetes Authors.

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

	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApintoRouters implements ApintoRouterInterface
type FakeApintoRouters struct {
	Fake *FakeApintoV1
	ns   string
}

var apintoroutersResource = schema.GroupVersionResource{Group: "apinto.com", Version: "v1", Resource: "apintorouters"}

var apintoroutersKind = schema.GroupVersionKind{Group: "apinto.com", Version: "v1", Kind: "ApintoRouter"}

// Get takes name of the apintoRouter, and returns the corresponding apintoRouter object, and an error if there is any.
func (c *FakeApintoRouters) Get(ctx context.Context, name string, options v1.GetOptions) (result *apintov1.ApintoRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apintoroutersResource, c.ns, name), &apintov1.ApintoRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoRouter), err
}

// List takes label and field selectors, and returns the list of ApintoRouters that match those selectors.
func (c *FakeApintoRouters) List(ctx context.Context, opts v1.ListOptions) (result *apintov1.ApintoRouterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apintoroutersResource, apintoroutersKind, c.ns, opts), &apintov1.ApintoRouterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apintov1.ApintoRouterList{ListMeta: obj.(*apintov1.ApintoRouterList).ListMeta}
	for _, item := range obj.(*apintov1.ApintoRouterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apintoRouters.
func (c *FakeApintoRouters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apintoroutersResource, c.ns, opts))

}

// Create takes the representation of a apintoRouter and creates it.  Returns the server's representation of the apintoRouter, and an error, if there is any.
func (c *FakeApintoRouters) Create(ctx context.Context, apintoRouter *apintov1.ApintoRouter, opts v1.CreateOptions) (result *apintov1.ApintoRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apintoroutersResource, c.ns, apintoRouter), &apintov1.ApintoRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoRouter), err
}

// Update takes the representation of a apintoRouter and updates it. Returns the server's representation of the apintoRouter, and an error, if there is any.
func (c *FakeApintoRouters) Update(ctx context.Context, apintoRouter *apintov1.ApintoRouter, opts v1.UpdateOptions) (result *apintov1.ApintoRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apintoroutersResource, c.ns, apintoRouter), &apintov1.ApintoRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoRouter), err
}

// Delete takes name of the apintoRouter and deletes it. Returns an error if one occurs.
func (c *FakeApintoRouters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apintoroutersResource, c.ns, name, opts), &apintov1.ApintoRouter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApintoRouters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apintoroutersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apintov1.ApintoRouterList{})
	return err
}

// Patch applies the patch and returns the patched apintoRouter.
func (c *FakeApintoRouters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apintov1.ApintoRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apintoroutersResource, c.ns, name, pt, data, subresources...), &apintov1.ApintoRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoRouter), err
}
