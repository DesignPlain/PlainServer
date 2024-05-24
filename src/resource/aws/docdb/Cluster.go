package docdb

type Cluster struct {
	// The database engine version. Updating this argument results in an outage.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Username for the master DB user.
	MasterUsername string `json:"masterUsername,omitempty" yaml:"masterUsername,omitempty"`

	// A map of tags to assign to the DB cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The days to retain backups for. Default `1`
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" yaml:"backupRetentionPeriod,omitempty"`

	/*
	   Password for the master DB user. Note that this may
	   show up in logs, and it will be stored in the state file. Please refer to the DocumentDB Naming Constraints.
	*/
	MasterPassword string `json:"masterPassword,omitempty" yaml:"masterPassword,omitempty"`

	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" yaml:"skipFinalSnapshot,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is `false`.
	StorageEncrypted bool `json:"storageEncrypted,omitempty" yaml:"storageEncrypted,omitempty"`

	/*
	   List of log types to export to cloudwatch. If omitted, no logs will be exported.
	   The following log types are supported: `audit`, `profiler`.
	*/
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" yaml:"enabledCloudwatchLogsExports,omitempty"`

	/*
	   The name of your final DB snapshot
	   when this DB cluster is deleted. If omitted, no final snapshot will be
	   made.
	*/
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot. Automated snapshots --should not-- be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" yaml:"snapshotIdentifier,omitempty"`

	/*
	   A list of EC2 Availability Zones that
	   instances in the DB cluster can be created in.
	*/
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	/*
	   The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
	   Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	*/
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" yaml:"preferredBackupWindow,omitempty"`

	/*
	   List of VPC security groups to associate
	   with the Cluster
	*/
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// The cluster identifier. If omitted, the provider will assign a random, unique identifier.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifier`.
	ClusterIdentifierPrefix string `json:"clusterIdentifierPrefix,omitempty" yaml:"clusterIdentifierPrefix,omitempty"`

	// A DB subnet group to associate with this DB instance.
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" yaml:"dbSubnetGroupName,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// The storage type to associate with the DB cluster. Valid values: `standard`, `iopt1`.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// A value that indicates whether major version upgrades are allowed. Constraints: You must allow major version upgrades when specifying a value for the EngineVersion parameter that is a different major version than the DB cluster's current version.
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" yaml:"allowMajorVersionUpgrade,omitempty"`

	// A cluster parameter group to associate with the cluster.
	DbClusterParameterGroupName string `json:"dbClusterParameterGroupName,omitempty" yaml:"dbClusterParameterGroupName,omitempty"`

	// The global cluster identifier specified on `aws.docdb.GlobalCluster`.
	GlobalClusterIdentifier string `json:"globalClusterIdentifier,omitempty" yaml:"globalClusterIdentifier,omitempty"`

	/*
	   Specifies whether any cluster modifications
	   are applied immediately, or during the next maintenance window. Default is
	   `false`.
	*/
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// The ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The port on which the DB accepts connections
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// List of DocumentDB Instances that are a part of this cluster
	ClusterMembers []string `json:"clusterMembers,omitempty" yaml:"clusterMembers,omitempty"`

	// A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// The name of the database engine to be used for this DB cluster. Defaults to `docdb`. Valid values: `docdb`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`
}
