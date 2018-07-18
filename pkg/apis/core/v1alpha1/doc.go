
/*
   Copyright 2018 The HAWQ Team.
*/


// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/hawq-cn/apiserver-example/pkg/apis/core
// +k8s:defaulter-gen=TypeMeta
// +groupName=core.hawq.org
package v1alpha1 // import "github.com/hawq-cn/apiserver-example/pkg/apis/core/v1alpha1"

