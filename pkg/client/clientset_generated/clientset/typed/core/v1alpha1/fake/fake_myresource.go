/*
   Copyright 2018 The HAWQ Team.
*/
package fake

import (
	v1alpha1 "github.com/hawq-cn/apiserver-example/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyResources implements MyResourceInterface
type FakeMyResources struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var myresourcesResource = schema.GroupVersionResource{Group: "core.hawq.org", Version: "v1alpha1", Resource: "myresources"}

var myresourcesKind = schema.GroupVersionKind{Group: "core.hawq.org", Version: "v1alpha1", Kind: "MyResource"}

// Get takes name of the myResource, and returns the corresponding myResource object, and an error if there is any.
func (c *FakeMyResources) Get(name string, options v1.GetOptions) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(myresourcesResource, c.ns, name), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}

// List takes label and field selectors, and returns the list of MyResources that match those selectors.
func (c *FakeMyResources) List(opts v1.ListOptions) (result *v1alpha1.MyResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(myresourcesResource, myresourcesKind, c.ns, opts), &v1alpha1.MyResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MyResourceList{}
	for _, item := range obj.(*v1alpha1.MyResourceList).Items {
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
func (c *FakeMyResources) Create(myResource *v1alpha1.MyResource) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(myresourcesResource, c.ns, myResource), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}

// Update takes the representation of a myResource and updates it. Returns the server's representation of the myResource, and an error, if there is any.
func (c *FakeMyResources) Update(myResource *v1alpha1.MyResource) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(myresourcesResource, c.ns, myResource), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMyResources) UpdateStatus(myResource *v1alpha1.MyResource) (*v1alpha1.MyResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(myresourcesResource, "status", c.ns, myResource), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}

// Delete takes name of the myResource and deletes it. Returns an error if one occurs.
func (c *FakeMyResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(myresourcesResource, c.ns, name), &v1alpha1.MyResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(myresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MyResourceList{})
	return err
}

// Patch applies the patch and returns the patched myResource.
func (c *FakeMyResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(myresourcesResource, c.ns, name, data, subresources...), &v1alpha1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MyResource), err
}
