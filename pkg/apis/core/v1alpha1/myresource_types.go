
/*
   Copyright 2018 The HAWQ Team.
*/


package v1alpha1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/hawq-cn/apiserver-example/pkg/apis/core"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MyResource
// +k8s:openapi-gen=true
// +resource:path=myresources,strategy=MyResourceStrategy
type MyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyResourceSpec   `json:"spec,omitempty"`
	Status MyResourceStatus `json:"status,omitempty"`
}

// MyResourceSpec defines the desired state of MyResource
type MyResourceSpec struct {
}

// MyResourceStatus defines the observed state of MyResource
type MyResourceStatus struct {
}

// Validate checks that an instance of MyResource is well formed
func (MyResourceStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*core.MyResource)
	log.Printf("Validating fields for MyResource %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default MyResource field values
func (MyResourceSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*MyResource)
	// set default field values here
	log.Printf("Defaulting fields for MyResource %s\n", obj.Name)
}
