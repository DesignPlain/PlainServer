package drs

import types "libds/aws/types"

type ReplicationConfigurationTemplate struct {
	// Configuration block for Point in time (PIT) policy to manage snapshots taken during replication. See below.
	PitPolicies []types.Drs_ReplicationConfigurationTemplatePitPolicy `json:"pitPolicies,omitempty" yaml:"pitPolicies,omitempty"`

	// Instance type to be used for the replication server.
	ReplicationServerInstanceType string `json:"replicationServerInstanceType,omitempty" yaml:"replicationServerInstanceType,omitempty"`

	// Security group IDs that will be used by the replication server.
	ReplicationServersSecurityGroupsIds []string `json:"replicationServersSecurityGroupsIds,omitempty" yaml:"replicationServersSecurityGroupsIds,omitempty"`

	// Subnet to be used by the replication staging area.
	StagingAreaSubnetId string `json:"stagingAreaSubnetId,omitempty" yaml:"stagingAreaSubnetId,omitempty"`

	// Set of tags to be associated with the Replication Configuration Template resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Drs_ReplicationConfigurationTemplateTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Whether to allow the AWS replication agent to automatically replicate newly added disks.
	AutoReplicateNewDisks bool `json:"autoReplicateNewDisks,omitempty" yaml:"autoReplicateNewDisks,omitempty"`

	// Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.
	BandwidthThrottling int `json:"bandwidthThrottling,omitempty" yaml:"bandwidthThrottling,omitempty"`

	// Type of EBS encryption to be used during replication. Valid values are `DEFAULT` and `CUSTOM`.
	EbsEncryption string `json:"ebsEncryption,omitempty" yaml:"ebsEncryption,omitempty"`

	// ARN of the EBS encryption key to be used during replication.
	EbsEncryptionKeyArn string `json:"ebsEncryptionKeyArn,omitempty" yaml:"ebsEncryptionKeyArn,omitempty"`

	// Set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	StagingAreaTags map[string]string `json:"stagingAreaTags,omitempty" yaml:"stagingAreaTags,omitempty"`

	// Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.
	AssociateDefaultSecurityGroup bool `json:"associateDefaultSecurityGroup,omitempty" yaml:"associateDefaultSecurityGroup,omitempty"`

	// Data plane routing mechanism that will be used for replication. Valid values are `PUBLIC_IP` and `PRIVATE_IP`.
	DataPlaneRouting string `json:"dataPlaneRouting,omitempty" yaml:"dataPlaneRouting,omitempty"`

	// Staging Disk EBS volume type to be used during replication. Valid values are `GP2`, `GP3`, `ST1`, or `AUTO`.
	DefaultLargeStagingDiskType string `json:"defaultLargeStagingDiskType,omitempty" yaml:"defaultLargeStagingDiskType,omitempty"`

	/*
	   Whether to use a dedicated Replication Server in the replication staging area.

	   The following arguments are optional:
	*/
	UseDedicatedReplicationServer bool `json:"useDedicatedReplicationServer,omitempty" yaml:"useDedicatedReplicationServer,omitempty"`

	// Whether to create a Public IP for the Recovery Instance by default.
	CreatePublicIp bool `json:"createPublicIp,omitempty" yaml:"createPublicIp,omitempty"`
}
