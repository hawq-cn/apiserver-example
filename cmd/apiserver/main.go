
/*
   Copyright 2018 The HAWQ Team.
*/


package main

import (
	// Make sure dep tools picks up these dependencies
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "github.com/go-openapi/loads"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/cmd/server"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth

	"github.com/hawq-cn/apiserver-example/pkg/apis"
	"github.com/hawq-cn/apiserver-example/pkg/openapi"
)

func main() {
	version := "v0"
	server.StartApiServer("/registry/hawq.org", apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions, "Api", version)
}
