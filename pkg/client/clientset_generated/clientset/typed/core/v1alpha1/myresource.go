/*
   Copyright 2018 The HAWQ Team.
*/
package v1alpha1

import (
	v1alpha1 "github.com/hawq-cn/apiserver-example/pkg/apis/core/v1alpha1"
	scheme "github.com/hawq-cn/apiserver-example/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MyResourcesGetter has a method to return a MyResourceInterface.
// A group's client should implement this interface.
type MyResourcesGetter interface {
	MyResources(namespace string) MyResourceInterface
}

// MyResourceInterface has methods to work with MyResource resources.
type MyResourceInterface interface {
	Create(*v1alpha1.MyResource) (*v1alpha1.MyResource, error)
	Update(*v1alpha1.MyResource) (*v1alpha1.MyResource, error)
	UpdateStatus(*v1alpha1.MyResource) (*v1alpha1.MyResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MyResource, error)
	List(opts v1.ListOptions) (*v1alpha1.MyResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MyResource, err error)
	MyResourceExpansion
}

// myResources implements MyResourceInterface
type myResources struct {
	client rest.Interface
	ns     string
}

// newMyResources returns a MyResources
func newMyResources(c *CoreV1alpha1Client, namespace string) *myResources {
	return &myResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the myResource, and returns the corresponding myResource object, and an error if there is any.
func (c *myResources) Get(name string, options v1.GetOptions) (result *v1alpha1.MyResource, err error) {
	result = &v1alpha1.MyResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MyResources that match those selectors.
func (c *myResources) List(opts v1.ListOptions) (result *v1alpha1.MyResourceList, err error) {
	result = &v1alpha1.MyResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested myResources.
func (c *myResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a myResource and creates it.  Returns the server's representation of the myResource, and an error, if there is any.
func (c *myResources) Create(myResource *v1alpha1.MyResource) (result *v1alpha1.MyResource, err error) {
	result = &v1alpha1.MyResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("myresources").
		Body(myResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a myResource and updates it. Returns the server's representation of the myResource, and an error, if there is any.
func (c *myResources) Update(myResource *v1alpha1.MyResource) (result *v1alpha1.MyResource, err error) {
	result = &v1alpha1.MyResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("myresources").
		Name(myResource.Name).
		Body(myResource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *myResources) UpdateStatus(myResource *v1alpha1.MyResource) (result *v1alpha1.MyResource, err error) {
	result = &v1alpha1.MyResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("myresources").
		Name(myResource.Name).
		SubResource("status").
		Body(myResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the myResource and deletes it. Returns an error if one occurs.
func (c *myResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *myResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("myresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched myResource.
func (c *myResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MyResource, err error) {
	result = &v1alpha1.MyResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("myresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
