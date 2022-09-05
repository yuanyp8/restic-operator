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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	ResourceKindBackupSession     = "BackupSession"
	ResourceSingularBackupSession = "backupSession"
	ResourcePluralBackupSession   = "backupsessions"
)

// BackupSessionSpec defines the desired state of BackupSession
type BackupSessionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Invoker refers to the BackupConfiguration or backupBatch being used to invoke this backup session
	Invoker BackupInvokerRef `json:"invoker,omitempty"`
}

//+kubebuilder:validation:Enum=Pending;Skipped;Running;Succeeded;Failed;Unknown
type BackupSessionPhase string

const (
	BackupSessionPending   BackupSessionPhase = "Pending"
	BackupSessionSkipped   BackupSessionPhase = "Skipped"
	BackupSessionRunning   BackupSessionPhase = "Running"
	BackupSessionSucceeded BackupSessionPhase = "Succeeded"
	BackupSessionFailed    BackupSessionPhase = "Failed"
	BackupSessionUnknown   BackupSessionPhase = "Unknown"
)

//+kubebuilder:validation:Enum=Succeeded;Failed
type HostBackupPhase string

const (
	HostBackupSucceeded HostBackupPhase = "Succeeded"
	HostBackupFailed    HostBackupPhase = "Failed"
)

// +kubebuilder:validation:Enum=Pending;Succeeded;Running;Failed

type TargetPhase string

const (
	TargetBackupPending   TargetPhase = "Pending"
	TargetBackupSucceeded TargetPhase = "Succeeded"
	TargetBackupRunning   TargetPhase = "Running"
	TargetBackupFailed    TargetPhase = "Failed"
)

type BackupTargetStatus struct {
	Ref TargetRef `json:"ref,omitempty"`
}

// BackupSessionStatus defines the observed state of BackupSession
type BackupSessionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase           BackupSessionPhase `json:"phase,omitempty"`
	SessionDuration string             `json:"sessionDuration,omitempty"`
	//+optional
	Targets         []BackupTargetStatus `json:"targets,omitempty"`
	Conditions      []kmapi.Condition    `json:"conditions,omitempty"`
	SessionDeadline metav1.Time          `json:"sessionDeadline,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:path=backupsessions,singular=backupsession,categories={restic-operator,metavarse.com,all}
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Invoker-Type",type="string",JSONPath=".spec.invoker.kind"
//+kubebuilder:printcolumn:name="Invoker-Name",type="string",JSONPath=".spec.invoker.name"
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="Duration",type="string",JSONPath=".status.sessionDuration"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// BackupSession is the Schema for the backupsessions API
type BackupSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSessionSpec   `json:"spec,omitempty"`
	Status BackupSessionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BackupSessionList contains a list of BackupSession
type BackupSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupSession{}, &BackupSessionList{})
}
