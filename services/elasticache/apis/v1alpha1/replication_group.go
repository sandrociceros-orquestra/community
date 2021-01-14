// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReplicationGroupSpec defines the desired state of ReplicationGroup
type ReplicationGroupSpec struct {
	AtRestEncryptionEnabled    *bool                     `json:"atRestEncryptionEnabled,omitempty"`
	AuthToken                  *string                   `json:"authToken,omitempty"`
	AutoMinorVersionUpgrade    *bool                     `json:"autoMinorVersionUpgrade,omitempty"`
	AutomaticFailoverEnabled   *bool                     `json:"automaticFailoverEnabled,omitempty"`
	CacheNodeType              *string                   `json:"cacheNodeType,omitempty"`
	CacheParameterGroupName    *string                   `json:"cacheParameterGroupName,omitempty"`
	CacheSecurityGroupNames    []*string                 `json:"cacheSecurityGroupNames,omitempty"`
	CacheSubnetGroupName       *string                   `json:"cacheSubnetGroupName,omitempty"`
	Engine                     *string                   `json:"engine,omitempty"`
	EngineVersion              *string                   `json:"engineVersion,omitempty"`
	GlobalReplicationGroupID   *string                   `json:"globalReplicationGroupID,omitempty"`
	KMSKeyID                   *string                   `json:"kmsKeyID,omitempty"`
	MultiAZEnabled             *bool                     `json:"multiAZEnabled,omitempty"`
	NodeGroupConfiguration     []*NodeGroupConfiguration `json:"nodeGroupConfiguration,omitempty"`
	NotificationTopicARN       *string                   `json:"notificationTopicARN,omitempty"`
	NumCacheClusters           *int64                    `json:"numCacheClusters,omitempty"`
	NumNodeGroups              *int64                    `json:"numNodeGroups,omitempty"`
	Port                       *int64                    `json:"port,omitempty"`
	PreferredCacheClusterAZs   []*string                 `json:"preferredCacheClusterAZs,omitempty"`
	PreferredMaintenanceWindow *string                   `json:"preferredMaintenanceWindow,omitempty"`
	PrimaryClusterID           *string                   `json:"primaryClusterID,omitempty"`
	ReplicasPerNodeGroup       *int64                    `json:"replicasPerNodeGroup,omitempty"`
	// +kubebuilder:validation:Required
	ReplicationGroupDescription *string `json:"replicationGroupDescription"`
	// +kubebuilder:validation:Required
	ReplicationGroupID       *string   `json:"replicationGroupID"`
	SecurityGroupIDs         []*string `json:"securityGroupIDs,omitempty"`
	SnapshotARNs             []*string `json:"snapshotARNs,omitempty"`
	SnapshotName             *string   `json:"snapshotName,omitempty"`
	SnapshotRetentionLimit   *int64    `json:"snapshotRetentionLimit,omitempty"`
	SnapshotWindow           *string   `json:"snapshotWindow,omitempty"`
	Tags                     []*Tag    `json:"tags,omitempty"`
	TransitEncryptionEnabled *bool     `json:"transitEncryptionEnabled,omitempty"`
	UserGroupIDs             []*string `json:"userGroupIDs,omitempty"`
}

// ReplicationGroupStatus defines the observed state of ReplicationGroup
type ReplicationGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions                    []*ackv1alpha1.Condition               `json:"conditions"`
	AllowedScaleDownModifications []*string                              `json:"allowedScaleDownModifications,omitempty"`
	AllowedScaleUpModifications   []*string                              `json:"allowedScaleUpModifications,omitempty"`
	AuthTokenEnabled              *bool                                  `json:"authTokenEnabled,omitempty"`
	AuthTokenLastModifiedDate     *metav1.Time                           `json:"authTokenLastModifiedDate,omitempty"`
	AutomaticFailover             *string                                `json:"automaticFailover,omitempty"`
	ClusterEnabled                *bool                                  `json:"clusterEnabled,omitempty"`
	ConfigurationEndpoint         *Endpoint                              `json:"configurationEndpoint,omitempty"`
	Description                   *string                                `json:"description,omitempty"`
	Events                        []*Event                               `json:"events,omitempty"`
	GlobalReplicationGroupInfo    *GlobalReplicationGroupInfo            `json:"globalReplicationGroupInfo,omitempty"`
	MemberClusters                []*string                              `json:"memberClusters,omitempty"`
	MemberClustersOutpostARNs     []*string                              `json:"memberClustersOutpostARNs,omitempty"`
	MultiAZ                       *string                                `json:"multiAZ,omitempty"`
	NodeGroups                    []*NodeGroup                           `json:"nodeGroups,omitempty"`
	PendingModifiedValues         *ReplicationGroupPendingModifiedValues `json:"pendingModifiedValues,omitempty"`
	SnapshottingClusterID         *string                                `json:"snapshottingClusterID,omitempty"`
	Status                        *string                                `json:"status,omitempty"`
}

// ReplicationGroup is the Schema for the ReplicationGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ReplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationGroupSpec   `json:"spec,omitempty"`
	Status            ReplicationGroupStatus `json:"status,omitempty"`
}

// ReplicationGroupList contains a list of ReplicationGroup
// +kubebuilder:object:root=true
type ReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReplicationGroup{}, &ReplicationGroupList{})
}
