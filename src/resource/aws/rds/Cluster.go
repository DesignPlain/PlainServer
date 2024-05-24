package rds

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifier`.
	ClusterIdentifierPrefix string `json:"clusterIdentifierPrefix,omitempty" yaml:"clusterIdentifierPrefix,omitempty"`

	// The name of the IAM role to be used when making API calls to the Directory Service.
	DomainIamRoleName string `json:"domainIamRoleName,omitempty" yaml:"domainIamRoleName,omitempty"`

	// Database engine mode. Valid values: `global` (only valid for Aurora MySQL 1.21 and earlier), `parallelquery`, `provisioned`, `serverless`. Defaults to: `provisioned`. See the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/aurora-serverless.html) for limitations when using `serverless`.
	EngineMode string `json:"engineMode,omitempty" yaml:"engineMode,omitempty"`

	// List of ARNs for the IAM roles to associate to the RDS Cluster.
	IamRoles []string `json:"iamRoles,omitempty" yaml:"iamRoles,omitempty"`

	// Network type of the cluster. Valid values: `IPV4`, `DUAL`.
	NetworkType string `json:"networkType,omitempty" yaml:"networkType,omitempty"`

	// Weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	//
	S3Import types.Rds_ClusterS3Import `json:"s3Import,omitempty" yaml:"s3Import,omitempty"`

	/*
	   List of EC2 Availability Zones for the DB cluster storage where DB cluster instances can be created.
	   RDS automatically assigns 3 AZs if less than 3 AZs are configured, which will show as a difference requiring resource recreation next pulumi up.
	   We recommend specifying 3 AZs or using the `lifecycle` configuration block `ignore_changes` argument if necessary.
	   A maximum of 3 AZs can be configured.
	*/
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`. See [Amazon RDS Documentation for more information.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// Name for an automatically created database on cluster creation. There are different naming restrictions per database engine: [RDS Naming Constraints](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints)
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Enable HTTP endpoint (data API). Only valid when `engine_mode` is set to `serverless`.
	EnableHttpEndpoint bool `json:"enableHttpEndpoint,omitempty" yaml:"enableHttpEndpoint,omitempty"`

	// Name of your final DB snapshot when this DB cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// Global cluster identifier specified on `aws.rds.GlobalCluster`.
	GlobalClusterIdentifier string `json:"globalClusterIdentifier,omitempty" yaml:"globalClusterIdentifier,omitempty"`

	// ARN of a source DB cluster or DB instance if this DB cluster is to be created as a Read Replica. If DB Cluster is part of a Global Cluster, use the `lifecycle` configuration block `ignore_changes` argument to prevent this provider from showing differences for this argument instead of configuring this value.
	ReplicationSourceIdentifier string `json:"replicationSourceIdentifier,omitempty" yaml:"replicationSourceIdentifier,omitempty"`

	// Enable to allow major engine version upgrades when changing engine versions. Defaults to `false`.
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" yaml:"allowMajorVersionUpgrade,omitempty"`

	// Nested attribute with scaling properties. Only valid when `engine_mode` is set to `serverless`. More details below.
	ScalingConfiguration types.Rds_ClusterScalingConfiguration `json:"scalingConfiguration,omitempty" yaml:"scalingConfiguration,omitempty"`

	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" yaml:"skipFinalSnapshot,omitempty"`

	// List of VPC security groups to associate with the Cluster
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// Nested attribute for [point in time restore](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PIT.html). More details below.
	RestoreToPointInTime types.Rds_ClusterRestoreToPointInTime `json:"restoreToPointInTime,omitempty" yaml:"restoreToPointInTime,omitempty"`

	/*
	   DB subnet group to associate with this DB cluster.
	   --NOTE:-- This must match the `db_subnet_group_name` specified on every `aws.rds.ClusterInstance` in the cluster.
	*/
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" yaml:"dbSubnetGroupName,omitempty"`

	/*
	   If the DB cluster should have deletion protection enabled.
	   The database can't be deleted when this value is set to `true`.
	   The default is `false`.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// Amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster. For information about valid Iops values, see [Amazon RDS Provisioned IOPS storage to improve performance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the Amazon RDS User Guide. (This setting is required to create a Multi-AZ DB cluster). Must be a multiple between .5 and 50 of the storage amount for the DB cluster.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN. If not specified, the default KMS key for your Amazon Web Services account is used.
	MasterUserSecretKmsKeyId string `json:"masterUserSecretKmsKeyId,omitempty" yaml:"masterUserSecretKmsKeyId,omitempty"`

	// The source region for an encrypted replica DB cluster.
	SourceRegion string `json:"sourceRegion,omitempty" yaml:"sourceRegion,omitempty"`

	// A cluster parameter group to associate with the cluster.
	DbClusterParameterGroupName string `json:"dbClusterParameterGroupName,omitempty" yaml:"dbClusterParameterGroupName,omitempty"`

	// Name of the database engine to be used for this DB cluster. Valid Values: `aurora-mysql`, `aurora-postgresql`, `mysql`, `postgres`. (Note that `mysql` and `postgres` are Multi-AZ RDS clusters).
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Database engine version. Updating this argument results in an outage. See the [Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) and [Aurora Postgres](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.html) documentation for your configured engine to determine this value, or by running `aws rds describe-db-engine-versions`. For example with Aurora MySQL 2, a potential value for this argument is `5.7.mysql_aurora.2.03.2`. The value can contain a partial version where supported by the API. The actual engine version used is returned in the attribute `engine_version_actual`, , see Attribute Reference below.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Password for the master DB user. Note that this may show up in logs, and it will be stored in the state file. Please refer to the [RDS Naming Constraints](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints). Cannot be set if `manage_master_user_password` is set to `true`.
	MasterPassword string `json:"masterPassword,omitempty" yaml:"masterPassword,omitempty"`

	// Username for the master DB user. Please refer to the [RDS Naming Constraints](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints). This argument does not support in-place updates and cannot be changed during a restore from snapshot.
	MasterUsername string `json:"masterUsername,omitempty" yaml:"masterUsername,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot. Conflicts with `global_cluster_identifier`. Clusters cannot be restored from snapshot --and-- joined to an existing global cluster in a single operation. See the [AWS documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-getting-started.html#aurora-global-database.use-snapshot) or the Global Cluster Restored From Snapshot example for instructions on building a global cluster starting with a snapshot.
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" yaml:"snapshotIdentifier,omitempty"`

	// (Forces new for Multi-AZ DB clusters) Specifies the storage type to be associated with the DB cluster. For Aurora DB clusters, `storage_type` modifications can be done in-place. For Multi-AZ DB Clusters, the `iops` argument must also be set. Valid values are: `""`, `aurora-iopt1` (Aurora DB Clusters); `io1` (Multi-AZ DB Clusters). Default: `""` (Aurora DB Clusters); `io1` (Multi-AZ DB Clusters).
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// A map of tags to assign to the DB cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Days to retain backups for. Default `1`
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" yaml:"backupRetentionPeriod,omitempty"`

	// Nested attribute with scaling properties for ServerlessV2. Only valid when `engine_mode` is set to `provisioned`. More details below.
	Serverlessv2ScalingConfiguration types.Rds_ClusterServerlessv2ScalingConfiguration `json:"serverlessv2ScalingConfiguration,omitempty" yaml:"serverlessv2ScalingConfiguration,omitempty"`

	// Whether cluster should forward writes to an associated global cluster. Applied to secondary clusters to enable them to forward writes to an `aws.rds.GlobalCluster`'s primary cluster. See the [Aurora Userguide documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-write-forwarding.html) for more information.
	EnableGlobalWriteForwarding bool `json:"enableGlobalWriteForwarding,omitempty" yaml:"enableGlobalWriteForwarding,omitempty"`

	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example `db.m6g.xlarge`. Not all DB instance classes are available in all AWS Regions, or for all database engines. For the full list of DB instance classes and availability for your engine, see [DB instance class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide.
	DbClusterInstanceClass string `json:"dbClusterInstanceClass,omitempty" yaml:"dbClusterInstanceClass,omitempty"`

	// Specifies whether to remove automated backups immediately after the DB cluster is deleted. Default is `true`.
	DeleteAutomatedBackups bool `json:"deleteAutomatedBackups,omitempty" yaml:"deleteAutomatedBackups,omitempty"`

	// Set of log types to export to cloudwatch. If omitted, no logs will be exported. The following log types are supported: `audit`, `error`, `general`, `slowquery`, `postgresql` (PostgreSQL).
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" yaml:"enabledCloudwatchLogsExports,omitempty"`

	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) for availability and limitations.
	IamDatabaseAuthenticationEnabled bool `json:"iamDatabaseAuthenticationEnabled,omitempty" yaml:"iamDatabaseAuthenticationEnabled,omitempty"`

	// Port on which the DB accepts connections
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" yaml:"preferredBackupWindow,omitempty"`

	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.
	AllocatedStorage int `json:"allocatedStorage,omitempty" yaml:"allocatedStorage,omitempty"`

	// List of RDS Instances that are a part of this cluster
	ClusterMembers []string `json:"clusterMembers,omitempty" yaml:"clusterMembers,omitempty"`

	// Copy all Cluster `tags` to snapshots. Default is `false`.
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" yaml:"copyTagsToSnapshot,omitempty"`

	// Instance parameter group to associate with all instances of the DB cluster. The `db_instance_parameter_group_name` parameter is only valid in combination with the `allow_major_version_upgrade` parameter.
	DbInstanceParameterGroupName string `json:"dbInstanceParameterGroupName,omitempty" yaml:"dbInstanceParameterGroupName,omitempty"`

	// For use with RDS Custom.
	DbSystemId string `json:"dbSystemId,omitempty" yaml:"dbSystemId,omitempty"`

	// The ID of the Directory Service Active Directory domain to create the cluster in.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// Set to true to allow RDS to manage the master user password in Secrets Manager. Cannot be set if `master_password` is provided.
	ManageMasterUserPassword bool `json:"manageMasterUserPassword,omitempty" yaml:"manageMasterUserPassword,omitempty"`

	// Specifies whether the DB cluster is encrypted. The default is `false` for `provisioned` `engine_mode` and `true` for `serverless` `engine_mode`. When restoring an unencrypted `snapshot_identifier`, the `kms_key_id` argument must be provided to encrypt the restored cluster. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted bool `json:"storageEncrypted,omitempty" yaml:"storageEncrypted,omitempty"`

	// Target backtrack window, in seconds. Only available for `aurora` and `aurora-mysql` engines currently. To disable backtracking, set this value to `0`. Defaults to `0`. Must be between `0` and `259200` (72 hours)
	BacktrackWindow int `json:"backtrackWindow,omitempty" yaml:"backtrackWindow,omitempty"`
}
