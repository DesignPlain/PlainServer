package redshift

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency. Can only be changed if `availability_zone_relocation_enabled` is `true`.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The specific revision number of the database in the cluster
	ClusterRevisionNumber string `json:"clusterRevisionNumber,omitempty" yaml:"clusterRevisionNumber,omitempty"`

	// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, `skip_final_snapshot` must be false.
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// The ARN of the snapshot from which to create the new cluster. Conflicts with `snapshot_identifier`.
	SnapshotArn string `json:"snapshotArn,omitempty" yaml:"snapshotArn,omitempty"`

	// ID of the KMS key used to encrypt the cluster admin credentials secret.
	MasterPasswordSecretKmsKeyId string `json:"masterPasswordSecretKmsKeyId,omitempty" yaml:"masterPasswordSecretKmsKeyId,omitempty"`

	// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is `true`.
	AllowVersionUpgrade bool `json:"allowVersionUpgrade,omitempty" yaml:"allowVersionUpgrade,omitempty"`

	// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	ClusterSubnetGroupName string `json:"clusterSubnetGroupName,omitempty" yaml:"clusterSubnetGroupName,omitempty"`

	/*
	   The name of the first database to be created when the cluster is created.
	   If you do not provide a name, Amazon Redshift will create a default database called `dev`.
	*/
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The Elastic IP (EIP) address for the cluster.
	ElasticIp string `json:"elasticIp,omitempty" yaml:"elasticIp,omitempty"`

	// If true , enhanced VPC routing is enabled.
	EnhancedVpcRouting bool `json:"enhancedVpcRouting,omitempty" yaml:"enhancedVpcRouting,omitempty"`

	// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	/*
	   Whether to use AWS SecretsManager to manage the cluster admin credentials.
	   Conflicts with `master_password`.
	   One of `master_password` or `manage_master_password` is required unless `snapshot_identifier` is provided.
	*/
	ManageMasterPassword bool `json:"manageMasterPassword,omitempty" yaml:"manageMasterPassword,omitempty"`

	// The name of the cluster the source snapshot was created from.
	SnapshotClusterIdentifier string `json:"snapshotClusterIdentifier,omitempty" yaml:"snapshotClusterIdentifier,omitempty"`

	// If true, the cluster can be relocated to another availabity zone, either automatically by AWS or when requested. Default is `false`. Available for use on clusters from the RA3 instance family.
	AvailabilityZoneRelocationEnabled bool `json:"availabilityZoneRelocationEnabled,omitempty" yaml:"availabilityZoneRelocationEnabled,omitempty"`

	/*
	   The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	   The version selected runs on all the nodes in the cluster.
	*/
	ClusterVersion string `json:"clusterVersion,omitempty" yaml:"clusterVersion,omitempty"`

	// If true , the data in the cluster is encrypted at rest.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
	AutomatedSnapshotRetentionPeriod int `json:"automatedSnapshotRetentionPeriod,omitempty" yaml:"automatedSnapshotRetentionPeriod,omitempty"`

	// The public key for the cluster
	ClusterPublicKey string `json:"clusterPublicKey,omitempty" yaml:"clusterPublicKey,omitempty"`

	// The connection endpoint
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// The node type to be provisioned for the cluster.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
	NumberOfNodes int `json:"numberOfNodes,omitempty" yaml:"numberOfNodes,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored.
	   No longer supported by the AWS API.
	   Always returns `auto`.
	*/
	AquaConfigurationStatus string `json:"aquaConfigurationStatus,omitempty" yaml:"aquaConfigurationStatus,omitempty"`

	// The name of the parameter group to be associated with this cluster.
	ClusterParameterGroupName string `json:"clusterParameterGroupName,omitempty" yaml:"clusterParameterGroupName,omitempty"`

	// Username for the master DB user.
	MasterUsername string `json:"masterUsername,omitempty" yaml:"masterUsername,omitempty"`

	/*
	   The port number on which the cluster accepts incoming connections. Valid values are between `1115` and `65535`.
	   The cluster is accessible only via the JDBC and ODBC connection strings.
	   Part of the connection string requires the port on which the cluster will listen for incoming connections.
	   Default port is `5439`.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// The Cluster Identifier. Must be a lower case string.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.
	DefaultIamRoleArn string `json:"defaultIamRoleArn,omitempty" yaml:"defaultIamRoleArn,omitempty"`

	// Logging, documented below.
	Logging types.Redshift_ClusterLogging `json:"logging,omitempty" yaml:"logging,omitempty"`

	// The cluster type to use. Either `single-node` or `multi-node`.
	ClusterType string `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`

	// The default number of days to retain a manual snapshot. If the value is -1, the snapshot is retained indefinitely. This setting doesn't change the retention period of existing snapshots. Valid values are between `-1` and `3653`. Default value is `-1`.
	ManualSnapshotRetentionPeriod int `json:"manualSnapshotRetentionPeriod,omitempty" yaml:"manualSnapshotRetentionPeriod,omitempty"`

	/*
	   Password for the master DB user.
	   Conflicts with `manage_master_password`.
	   One of `master_password` or `manage_master_password` is required unless `snapshot_identifier` is provided.
	   Note that this may show up in logs, and it will be stored in the state file.
	   Password must contain at least 8 characters and contain at least one uppercase letter, one lowercase letter, and one number.
	*/
	MasterPassword string `json:"masterPassword,omitempty" yaml:"masterPassword,omitempty"`

	/*
	   The weekly time range (in UTC) during which automated cluster maintenance can occur.
	   Format: ddd:hh24:mi-ddd:hh24:mi
	*/
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" yaml:"skipFinalSnapshot,omitempty"`

	// The name of the snapshot from which to create the new cluster.  Conflicts with `snapshot_arn`.
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" yaml:"snapshotIdentifier,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
	IamRoles []string `json:"iamRoles,omitempty" yaml:"iamRoles,omitempty"`

	// The name of the maintenance track for the restored cluster. When you take a snapshot, the snapshot inherits the MaintenanceTrack value from the cluster. The snapshot might be on a different track than the cluster that was the source for the snapshot. For example, suppose that you take a snapshot of  a cluster that is on the current track and then change the cluster to be on the trailing track. In this case, the snapshot and the source cluster are on different tracks. Default value is `current`.
	MaintenanceTrackName string `json:"maintenanceTrackName,omitempty" yaml:"maintenanceTrackName,omitempty"`

	// Specifies if the Redshift cluster is multi-AZ.
	MultiAz bool `json:"multiAz,omitempty" yaml:"multiAz,omitempty"`

	// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount string `json:"ownerAccount,omitempty" yaml:"ownerAccount,omitempty"`

	// If true, the cluster can be accessed from a public network. Default is `true`.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	// Configuration of automatic copy of snapshots from one region to another. Documented below.
	SnapshotCopy types.Redshift_ClusterSnapshotCopy `json:"snapshotCopy,omitempty" yaml:"snapshotCopy,omitempty"`
}
