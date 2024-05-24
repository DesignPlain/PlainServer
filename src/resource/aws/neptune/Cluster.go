package neptune

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifier`.
	ClusterIdentifierPrefix string `json:"clusterIdentifierPrefix,omitempty" yaml:"clusterIdentifierPrefix,omitempty"`

	// The ARN for the KMS encryption key. When specifying `kms_key_arn`, `storage_encrypted` needs to be set to true.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" yaml:"preferredBackupWindow,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier string `json:"replicationSourceIdentifier,omitempty" yaml:"replicationSourceIdentifier,omitempty"`

	// A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// The global cluster identifier specified on `aws.neptune.GlobalCluster`.
	GlobalClusterIdentifier string `json:"globalClusterIdentifier,omitempty" yaml:"globalClusterIdentifier,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// The port on which the Neptune accepts connections. Default is `8182`.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Storage type associated with the cluster `standard/iopt1`. Default: `standard`
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// A Neptune subnet group to associate with this Neptune instance.
	NeptuneSubnetGroupName string `json:"neptuneSubnetGroupName,omitempty" yaml:"neptuneSubnetGroupName,omitempty"`

	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" yaml:"skipFinalSnapshot,omitempty"`

	// List of VPC security groups to associate with the Cluster
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// The days to retain backups for. Default `1`
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" yaml:"backupRetentionPeriod,omitempty"`

	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit` and `slowquery`.
	EnableCloudwatchLogsExports []string `json:"enableCloudwatchLogsExports,omitempty" yaml:"enableCloudwatchLogsExports,omitempty"`

	// The database engine version.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	IamRoles []string `json:"iamRoles,omitempty" yaml:"iamRoles,omitempty"`

	// A cluster parameter group to associate with the cluster.
	NeptuneClusterParameterGroupName string `json:"neptuneClusterParameterGroupName,omitempty" yaml:"neptuneClusterParameterGroupName,omitempty"`

	// Specifies whether upgrades between different major versions are allowed. You must set it to `true` when providing an `engine_version` parameter that uses a different major version than the DB cluster's current version. Default is `false`.
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" yaml:"allowMajorVersionUpgrade,omitempty"`

	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" yaml:"copyTagsToSnapshot,omitempty"`

	// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// The name of the DB parameter group to apply to all instances of the DB cluster.
	NeptuneInstanceParameterGroupName string `json:"neptuneInstanceParameterGroupName,omitempty" yaml:"neptuneInstanceParameterGroupName,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot. Automated snapshots --should not-- be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" yaml:"snapshotIdentifier,omitempty"`

	// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
	StorageEncrypted bool `json:"storageEncrypted,omitempty" yaml:"storageEncrypted,omitempty"`

	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IamDatabaseAuthenticationEnabled bool `json:"iamDatabaseAuthenticationEnabled,omitempty" yaml:"iamDatabaseAuthenticationEnabled,omitempty"`

	// If set, create the Neptune cluster as a serverless one. See Serverless for example block attributes.
	ServerlessV2ScalingConfiguration types.Neptune_ClusterServerlessV2ScalingConfiguration `json:"serverlessV2ScalingConfiguration,omitempty" yaml:"serverlessV2ScalingConfiguration,omitempty"`
}
