
/*
   Copyright 2018 The HAWQ Team.
*/


package myresource_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/rest"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/test"

	"github.com/hawq-cn/apiserver-example/pkg/apis"
	"github.com/hawq-cn/apiserver-example/pkg/client/clientset_generated/clientset"
	"github.com/hawq-cn/apiserver-example/pkg/openapi"
	"github.com/hawq-cn/apiserver-example/pkg/controller/sharedinformers"
	"github.com/hawq-cn/apiserver-example/pkg/controller/myresource"
)

var testenv *test.TestEnvironment
var config *rest.Config
var cs *clientset.Clientset
var shutdown chan struct{}
var controller *myresource.MyResourceController
var si *sharedinformers.SharedInformers

func TestMyResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "MyResource Suite", []Reporter{test.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	testenv = test.NewTestEnvironment()
	config = testenv.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	cs = clientset.NewForConfigOrDie(config)

	shutdown = make(chan struct{})
	si = sharedinformers.NewSharedInformers(config, shutdown)
	controller = myresource.NewMyResourceController(config, si)
	controller.Run(shutdown)
})

var _ = AfterSuite(func() {
	close(shutdown)
	testenv.Stop()
})
