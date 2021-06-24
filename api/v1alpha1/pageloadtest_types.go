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

// PageLoadTestSpec defines the desired state of PageLoadTest
type PageLoadTestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PageLoad `json:",inline"`
}

// PageLoadTestStatus defines the observed state of PageLoadTest
type PageLoadTestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PageLoadTest is the Schema for the pageloadtest API
type PageLoadTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PageLoadTestSpec   `json:"spec,omitempty"`
	Status PageLoadTestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PageLoadTestList contains a list of PageLoadTest
type PageLoadTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PageLoadTest `json:"items"`
}

type PageLoad struct {
	TestID       int         `json:"testId,omitempty"`
	Type         string      `json:"type,omitempty"`
	Agents       []Agent     `json:"agents,omitempty"`
	HTTPInterval int         `json:"httpInterval,omitempty"`
	Interval     int         `json:"interval,omitempty"`
	URL          string      `json:"url,omitempty"`
	AlertRules   []AlertRule `json:"alertRules,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PageLoadTest{}, &PageLoadTestList{})
}
