/*
Copyright 2022.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=Starburst;Blue
type ISV string

const (
	Starburst ISV = "Starburst"
	Blue      ISV = "Blue"
)

// Loglevel set log levels of configured components
// +kubebuilder:validation:Enum=debug;info;warning
type LogLevel string

const (
	// Debug Log level
	Debug LogLevel = "debug"

	// Info Log level
	Info LogLevel = "info"

	// Warning Log level
	Warning LogLevel = "warning"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MSOSpec defines the desired state of MSO
type MSOSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	// +kubebuilder:default="info"
	LogLevel LogLevel `json:"logLevel,omitempty"`

	// Namespace to deploy mso
	// optional
	// +kubebuilder:default="monitoring-stack-operator"
	Namespace string `json:"namespace,omitempty"`

	// name of ISV to monitor, to determine which servicemonitors to apply
	// +optional
	// +kubebuilder:default="Starburst"
	ISV ISV `json:"isv,omitempty"`

	// Namespace of ISV to monitor
	IsvNamespace string `json:"isvNamespace,omitempty"`

	// URL to apply manifests
	// +optional
	ManifestUrl string `json:"manifestUrl,omitempty"`

	// Define resources requests and limits for Monitoring Stack Pods.
	// +optional
	// +kubebuilder:default={requests:{cpu: "100m", memory: "256M"}, limits:{memory: "512M", cpu: "500m"}}
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

// MSOStatus defines the observed state of MSO
type MSOStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// MSOs are the names of the mso pods
	MSOs []string `json:"msos"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MSO is the Schema for the msoes API
type MSO struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MSOSpec   `json:"spec,omitempty"`
	Status MSOStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MSOList contains a list of MSO
type MSOList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSO `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MSO{}, &MSOList{})
}
