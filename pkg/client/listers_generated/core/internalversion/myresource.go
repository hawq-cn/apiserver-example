/*
   Copyright 2018 The HAWQ Team.
*/

// This file was automatically generated by lister-gen

package internalversion

import (
	core "github.com/hawq-cn/apiserver-example/pkg/apis/core"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MyResourceLister helps list MyResources.
type MyResourceLister interface {
	// List lists all MyResources in the indexer.
	List(selector labels.Selector) (ret []*core.MyResource, err error)
	// MyResources returns an object that can list and get MyResources.
	MyResources(namespace string) MyResourceNamespaceLister
	MyResourceListerExpansion
}

// myResourceLister implements the MyResourceLister interface.
type myResourceLister struct {
	indexer cache.Indexer
}

// NewMyResourceLister returns a new MyResourceLister.
func NewMyResourceLister(indexer cache.Indexer) MyResourceLister {
	return &myResourceLister{indexer: indexer}
}

// List lists all MyResources in the indexer.
func (s *myResourceLister) List(selector labels.Selector) (ret []*core.MyResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*core.MyResource))
	})
	return ret, err
}

// MyResources returns an object that can list and get MyResources.
func (s *myResourceLister) MyResources(namespace string) MyResourceNamespaceLister {
	return myResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MyResourceNamespaceLister helps list and get MyResources.
type MyResourceNamespaceLister interface {
	// List lists all MyResources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*core.MyResource, err error)
	// Get retrieves the MyResource from the indexer for a given namespace and name.
	Get(name string) (*core.MyResource, error)
	MyResourceNamespaceListerExpansion
}

// myResourceNamespaceLister implements the MyResourceNamespaceLister
// interface.
type myResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MyResources in the indexer for a given namespace.
func (s myResourceNamespaceLister) List(selector labels.Selector) (ret []*core.MyResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*core.MyResource))
	})
	return ret, err
}

// Get retrieves the MyResource from the indexer for a given namespace and name.
func (s myResourceNamespaceLister) Get(name string) (*core.MyResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(core.Resource("myresource"), name)
	}
	return obj.(*core.MyResource), nil
}
