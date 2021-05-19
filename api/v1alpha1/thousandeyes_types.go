/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ThousandEyesSpec defines the desired state of ThousandEyes
type ThousandEyesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ThousandEyes. Edit thousandeyes_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ThousandEyesStatus defines the observed state of ThousandEyes
type ThousandEyesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ThousandEyes is the Schema for the thousandeyes API
type ThousandEyes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ThousandEyesSpec   `json:"spec,omitempty"`
	Status ThousandEyesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ThousandEyesList contains a list of ThousandEyes
type ThousandEyesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ThousandEyes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ThousandEyes{}, &ThousandEyesList{})
}
