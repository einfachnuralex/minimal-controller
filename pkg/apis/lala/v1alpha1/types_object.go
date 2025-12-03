package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=lalas,scope=Cluster

// Lala is the schema for the Lala API
type Lala struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LalaSpec   `json:"spec,omitempty"`
	Status LalaStatus `json:"status,omitempty"`
}

// LalaSpec defines the observed state of the LalaSpec
type LalaSpec struct {
	Name   string `json:"name"`
	Reason string `json:"reason,omitempty"`
}

// LalaStatus defines the observed state of Lala
type LalaStatus struct {
	// InstanceID is the ID of the instance of the operator
	InstanceID string `json:"instanceID,omitempty"`
}

// +kubebuilder:object:root=true
// LalaList contains a list of Lala
type LalaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lala `json:"items"`
}
