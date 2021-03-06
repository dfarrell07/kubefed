/*
Copyright 2018 The Kubernetes Authors.

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
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/kubefed/pkg/apis/core/common"
)

// KubefedClusterSpec defines the desired state of KubefedCluster
type KubefedClusterSpec struct {
	// The API endpoint of the member cluster. This can be a hostname,
	// hostname:port, IP or IP:port.
	APIEndpoint string `json:"apiEndpoint"`

	// CABundle contains the certificate authority information.
	// +optional
	CABundle []byte `json:"caBundle,omitempty"`

	// Name of the secret containing the token required to access the
	// member cluster. The secret needs to exist in the same namespace
	// as the control plane and should have a "token" key.
	SecretRef LocalSecretReference `json:"secretRef"`
}

// LocalSecretReference is a reference to a secret within the enclosing
// namespace.
type LocalSecretReference struct {
	// Name of a secret within the enclosing
	// namespace
	Name string `json:"name"`
}

// KubefedClusterStatus contains information about the current status of a
// cluster updated periodically by cluster controller.
type KubefedClusterStatus struct {
	// Conditions is an array of current cluster conditions.
	// +optional
	Conditions []ClusterCondition `json:"conditions,omitempty"`
	// Zones are the names of availability zones in which the nodes of the cluster exist, e.g. 'us-east1-a'.
	// +optional
	Zones []string `json:"zones,omitempty"`
	// Region is the name of the region in which all of the nodes in the cluster exist.  e.g. 'us-east1'.
	// +optional
	Region string `json:"region,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubefedCluster configures federation to be aware of a Kubernetes
// cluster and provides a Kubeconfig for federation to use to
// communicate with the cluster.
//
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=kubefedclusters
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name=ready,type=string,JSONPath=.status.conditions[?(@.type=='Ready')].status
// +kubebuilder:printcolumn:name=age,type=date,JSONPath=.metadata.creationTimestamp
type KubefedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubefedClusterSpec   `json:"spec,omitempty"`
	Status KubefedClusterStatus `json:"status,omitempty"`
}

// ClusterCondition describes current state of a cluster.
type ClusterCondition struct {
	// Type of cluster condition, Ready or Offline.
	Type common.ClusterConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status apiv1.ConditionStatus `json:"status"`
	// Last time the condition was checked.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubefedClusterList contains a list of KubefedCluster
type KubefedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubefedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubefedCluster{}, &KubefedClusterList{})
}
