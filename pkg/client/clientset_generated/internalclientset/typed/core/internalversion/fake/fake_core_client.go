/*
   Copyright 2018 The HAWQ Team.
*/
package fake

import (
	internalversion "github.com/hawq-cn/apiserver-example/pkg/client/clientset_generated/internalclientset/typed/core/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCore struct {
	*testing.Fake
}

func (c *FakeCore) MyResources(namespace string) internalversion.MyResourceInterface {
	return &FakeMyResources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCore) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
