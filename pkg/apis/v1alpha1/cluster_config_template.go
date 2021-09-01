// Copyright 2021 VMware
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +versionName=v1alpha1
// +groupName=carto.run
// +kubebuilder:object:generate=true

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

type ClusterConfigTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ConfigTemplateSpec   `json:"spec"`
	Status            ConfigTemplateStatus `json:"status,omitempty"`
}

type ConfigTemplateSpec struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	Template   runtime.RawExtension `json:"template"`
	ConfigPath string               `json:"configPath"`
	Params     DefaultParams        `json:"params,omitempty"`
}

type ConfigTemplateStatus struct {
}

// +kubebuilder:object:root=true

type ClusterConfigTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterConfigTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ClusterConfigTemplate{},
		&ClusterConfigTemplateList{},
	)
}
