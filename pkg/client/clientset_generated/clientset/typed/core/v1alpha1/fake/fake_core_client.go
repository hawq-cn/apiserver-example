/*
   Copyright 2018 The HAWQ Team.
*/
package fake

import (
	v1alpha1 "github.com/hawq-cn/apiserver-example/pkg/client/clientset_generated/clientset/typed/core/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCoreV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCoreV1alpha1) MyResources(namespace string) v1alpha1.MyResourceInterface {
	return &FakeMyResources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
