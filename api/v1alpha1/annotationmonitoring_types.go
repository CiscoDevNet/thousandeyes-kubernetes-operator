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

// AnnotationMonitoringSpec defines the desired state of AnnotationMonitoring
type AnnotationMonitoringSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

// AnnotationMonitoringStatus defines the observed state of AnnotationMonitoring
type AnnotationMonitoringStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AnnotationMonitoring is the Schema for the annotationmonitorings API
type AnnotationMonitoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnnotationMonitoringSpec   `json:"spec,omitempty"`
	Status AnnotationMonitoringStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AnnotationMonitoringList contains a list of AnnotationMonitoring
type AnnotationMonitoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnnotationMonitoring `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AnnotationMonitoring{}, &AnnotationMonitoringList{})
}
