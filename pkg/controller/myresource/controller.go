
/*
   Copyright 2018 The HAWQ Team.
*/


package myresource

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/hawq-cn/apiserver-example/pkg/apis/core/v1alpha1"
	"github.com/hawq-cn/apiserver-example/pkg/controller/sharedinformers"
	listers "github.com/hawq-cn/apiserver-example/pkg/client/listers_generated/core/v1alpha1"
)

// +controller:group=core,version=v1alpha1,kind=MyResource,resource=myresources
type MyResourceControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about MyResource
	lister listers.MyResourceLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *MyResourceControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing myresources labels
	c.lister = arguments.GetSharedInformers().Factory.Core().V1alpha1().MyResources().Lister()
}

// Reconcile handles enqueued messages
func (c *MyResourceControllerImpl) Reconcile(u *v1alpha1.MyResource) error {
	// Implement controller logic here
	log.Printf("Running reconcile MyResource for %s\n", u.Name)
	return nil
}

func (c *MyResourceControllerImpl) Get(namespace, name string) (*v1alpha1.MyResource, error) {
	return c.lister.MyResources(namespace).Get(name)
}
