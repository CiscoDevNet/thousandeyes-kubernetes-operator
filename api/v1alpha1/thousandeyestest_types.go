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

// ThousandEyesTestSpec defines the desired state of ThousandEyesTest
type ThousandEyesTestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	TestType string   `json:"testType"`
	Metadata Metadata `json:"metadata"`
}

// ThousandEyesTestStatus defines the observed state of ThousandEyesTest
type ThousandEyesTestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ThousandEyesTest is the Schema for the thousandeyestests API
type ThousandEyesTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ThousandEyesTestSpec   `json:"spec,omitempty"`
	Status ThousandEyesTestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ThousandEyesTestList contains a list of ThousandEyesTest
type ThousandEyesTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ThousandEyesTest `json:"items"`
}

type Metadata struct {
	TestName     string  `json:"testName,omitempty"`
	TestID       int     `json:"testId,omitempty"`
	URL          string  `json:"url"`
	HttpInterval int     `json:"httpInterval"`
	Interval     int     `json:"interval"`
	Agents       []Agent `json:"agents"`
}

type Agent struct {
	AgentID int `json:"agentId"`
}

func init() {
	SchemeBuilder.Register(&ThousandEyesTest{}, &ThousandEyesTestList{})
}
