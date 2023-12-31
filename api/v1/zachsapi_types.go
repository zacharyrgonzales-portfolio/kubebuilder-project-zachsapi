/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ZachsAPISpec struct {
	ZachsPhoneNumber string `json:"zachsPhoneNumber,omitempty"`
	ZachsAge string `json:"zachsAge,omitempty"`
}

// ZachsAPIStatus defines the observed state of ZachsAPI
type ZachsAPIStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ZachsAPI is the Schema for the zachsapis API
type ZachsAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZachsAPISpec   `json:"spec,omitempty"`
	Status ZachsAPIStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ZachsAPIList contains a list of ZachsAPI
type ZachsAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZachsAPI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZachsAPI{}, &ZachsAPIList{})
}
