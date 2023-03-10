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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigMapSpec defines the desired state of ConfigMap
type ConfigMapSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	HbaConf []HbaConfLine  `json:"hbaConf,omitempty"`
	PGConf  PostgresqlConf `json:"pgConf,omitempty"`
}

type HbaConfLine struct {
	// hba.conf connection settings - sample https://www.postgresql.org/docs/current/auth-pg-hba-conf.html
	Type     string `json:"type,omitempty"`
	Database string `json:"database,omitempty"`
	User     string `json:"user,omitempty"`
	Address  string `json:"address,omitempty"`
	IpMask   string `json:"ipMask,omitempty"`
	Method   string `json:"method,omitempty"`
}

type PostgresqlConf struct {
	// postgresql.conf server settings - sample https://github.com/postgres/postgres/blob/master/src/backend/utils/misc/postgresql.conf.sample
	DataDir        string `json:"dataDir,omitempty"`
	HbaFile        string `json:"hbaFile,omitempty"`
	Port           int    `json:"port,omitempty"`
	MaxConnections int    `json:"maxConnections,omitempty"`
}

// ConfigMapStatus defines the observed state of ConfigMap
type ConfigMapStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConfigMap is the Schema for the configmaps API
type ConfigMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigMapSpec   `json:"spec,omitempty"`
	Status ConfigMapStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConfigMapList contains a list of ConfigMap
type ConfigMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigMap{}, &ConfigMapList{})
}
