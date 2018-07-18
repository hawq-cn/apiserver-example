/*
   Copyright 2018 The HAWQ Team.
*/
package fake

import (
	core "github.com/hawq-cn/apiserver-example/pkg/apis/core"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyResources implements MyResourceInterface
type FakeMyResources struct {
	Fake *FakeCore
	ns   string
}

var myresourcesResource = schema.GroupVersionResource{Group: "core.hawq.org", Version: "", Resource: "myresources"}

var myresourcesKind = schema.GroupVersionKind{Group: "core.hawq.org", Version: "", Kind: "MyResource"}

// Get takes name of the myResource, and returns the corresponding myResource object, and an error if there is any.
func (c *FakeMyResources) Get(name string, options v1.GetOptions) (result *core.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(myresourcesResource, c.ns, name), &core.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.MyResource), err
}

// List takes label and field selectors, and returns the list of MyResources that match those selectors.
func (c *FakeMyResources) List(opts v1.ListOptions) (result *core.MyResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(myresourcesResource, myresourcesKind, c.ns, opts), &core.MyResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.MyResourceList{}
	for _, item := range obj.(*core.MyResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested myResources.
func (c *FakeMyResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(myresourcesResource, c.ns, opts))

}

// Create takes the representation of a myResource and creates it.  Returns the server's representation of the myResource, and an error, if there is any.
func (c *FakeMyResources) Create(myResource *core.MyResource) (result *core.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(myresourcesResource, c.ns, myResource), &core.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.MyResource), err
}

// Update takes the representation of a myResource and updates it. Returns the server's representation of the myResource, and an error, if there is any.
func (c *FakeMyResources) Update(myResource *core.MyResource) (result *core.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(myresourcesResource, c.ns, myResource), &core.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.MyResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMyResources) UpdateStatus(myResource *core.MyResource) (*core.MyResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(myresourcesResource, "status", c.ns, myResource), &core.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.MyResource), err
}

// Delete takes name of the myResource and deletes it. Returns an error if one occurs.
func (c *FakeMyResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(myresourcesResource, c.ns, name), &core.MyResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(myresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &core.MyResourceList{})
	return err
}

// Patch applies the patch and returns the patched myResource.
func (c *FakeMyResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(myresourcesResource, c.ns, name, data, subresources...), &core.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.MyResource), err
}
