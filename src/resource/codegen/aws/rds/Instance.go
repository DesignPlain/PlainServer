package rds

import types "libds/aws/types"

type Instance struct {
	// Use a dedicated log volume (DLV) for the DB instance. Requires Provisioned IOPS. See the [AWS documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.dlv) for more details.
	DedicatedLogVolume bool `json:"dedicatedLogVolume,omitempty" yaml:"dedicatedLogVolume,omitempty"`

	/*
	   The national character set is used in the NCHAR, NVARCHAR2, and NCLOB data types for Oracle instances. This can't be changed. See [Oracle Character Sets
	   Supported in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html).
	*/
	NcharCharacterSetName string `json:"ncharCharacterSetName,omitempty" yaml:"ncharCharacterSetName,omitempty"`

	// Restore from a Percona Xtrabackup in S3.  See [Importing Data into an Amazon RDS MySQL DB Instance](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.html)
	S3Import types.Rds_InstanceS3Import `json:"s3Import,omitempty" yaml:"s3Import,omitempty"`

	// When configured, the upper limit to which Amazon RDS can automatically scale the storage of the DB instance. Configuring this will automatically ignore differences to `allocated_storage`. Must be greater than or equal to `allocated_storage` or `0` to disable Storage Autoscaling.
	MaxAllocatedStorage int `json:"maxAllocatedStorage,omitempty" yaml:"maxAllocatedStorage,omitempty"`

	/*
	   The ARN for the IAM role that permits RDS
	   to send enhanced monitoring metrics to CloudWatch Logs. You can find more
	   information on the [AWS
	   Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
	   what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	*/
	MonitoringRoleArn string `json:"monitoringRoleArn,omitempty" yaml:"monitoringRoleArn,omitempty"`

	// The port on which the DB accepts connections.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   Specifies that this resource is a Replicate
	   database, and to use this value as the source database. This correlates to the
	   `identifier` of another Amazon RDS Database to replicate (if replicating within
	   a single region) or ARN of the Amazon RDS Database to replicate (if replicating
	   cross-region). Note that if you are
	   creating a cross-region replica of an encrypted database you will also need to
	   specify a `kms_key_id`. See [DB Instance Replication][instance-replication] and [Working with
	   PostgreSQL and MySQL Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)
	   for more information on using Replication.
	*/
	ReplicateSourceDb string `json:"replicateSourceDb,omitempty" yaml:"replicateSourceDb,omitempty"`

	/*
	   List of VPC security groups to
	   associate.
	*/
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// The name of the IAM role to be used when making API calls to the Directory Service. Conflicts with `domain_fqdn`, `domain_ou`, `domain_auth_secret_arn` and a `domain_dns_ips`.
	DomainIamRoleName string `json:"domainIamRoleName,omitempty" yaml:"domainIamRoleName,omitempty"`

	// Name of the DB parameter group to associate.
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	/*
	   (Required unless `manage_master_user_password` is set to true or unless a `snapshot_identifier` or `replicate_source_db`
	   is provided or `manage_master_user_password` is set.) Password for the master DB user. Note that this may show up in
	   logs, and it will be stored in the state file. Cannot be set if `manage_master_user_password` is set to `true`.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	/*
	   Determines whether a final DB snapshot is
	   created before the DB instance is deleted. If true is specified, no DBSnapshot
	   is created. If false is specified, a DB snapshot is created before the DB
	   instance is deleted, using the value from `final_snapshot_identifier`. Default
	   is `false`.
	*/
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" yaml:"skipFinalSnapshot,omitempty"`

	/*
	   (Required unless a `snapshot_identifier` or `replicate_source_db`
	   is provided) Username for the master DB user. Cannot be specified for a replica.
	*/
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	/*
	   The days to retain backups for.
	   Must be between `0` and `35`.
	   Default is `0`.
	   Must be greater than `0` if the database is used as a source for a [Read Replica][instance-replication],
	   uses low-downtime updates,
	   or will use [RDS Blue/Green deployments][blue-green].
	*/
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" yaml:"backupRetentionPeriod,omitempty"`

	/*
	   Indicates whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance. See [CoIP for RDS on Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html#rds-on-outposts.coip) for more information.

	   > --NOTE:-- Removing the `replicate_source_db` attribute from an existing RDS
	   Replicate database managed by the provider will promote the database to a fully
	   standalone database.
	*/
	CustomerOwnedIpEnabled bool `json:"customerOwnedIpEnabled,omitempty" yaml:"customerOwnedIpEnabled,omitempty"`

	// Set of log types to enable for exporting to CloudWatch logs. If omitted, no logs will be exported. For supported values, see the EnableCloudwatchLogsExports.member.N parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" yaml:"enabledCloudwatchLogsExports,omitempty"`

	/*
	   The name of your final DB snapshot
	   when this DB instance is deleted. Must be provided if `skip_final_snapshot` is
	   set to `false`. The value must begin with a letter, only contain alphanumeric characters and hyphens, and not end with a hyphen or contain two consecutive hyphens. Must not be provided when deleting a read replica.
	*/
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// The network type of the DB instance. Valid values: `IPV4`, `DUAL`.
	NetworkType string `json:"networkType,omitempty" yaml:"networkType,omitempty"`

	// The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	CustomIamInstanceProfile string `json:"customIamInstanceProfile,omitempty" yaml:"customIamInstanceProfile,omitempty"`

	// The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance. Note that this does not apply for Oracle or SQL Server engines. See the [AWS documentation](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/create-db-instance.html) for more details on what applies for those engines. If you are providing an Oracle db name, it needs to be in all upper case. Cannot be specified for a replica.
	DbName string `json:"dbName,omitempty" yaml:"dbName,omitempty"`

	// The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN. If not specified, the default KMS key for your Amazon Web Services account is used.
	MasterUserSecretKmsKeyId string `json:"masterUserSecretKmsKeyId,omitempty" yaml:"masterUserSecretKmsKeyId,omitempty"`

	// Whether to upgrade the storage file system configuration on the read replica. Can only be set with `replicate_source_db`.
	UpgradeStorageConfig bool `json:"upgradeStorageConfig,omitempty" yaml:"upgradeStorageConfig,omitempty"`

	// The allocated storage in gibibytes. If `max_allocated_storage` is configured, this argument represents the initial storage allocation and differences from the configuration will be ignored automatically when Storage Autoscaling occurs. If `replicate_source_db` is set, the value is ignored during the creation of the instance.
	AllocatedStorage int `json:"allocatedStorage,omitempty" yaml:"allocatedStorage,omitempty"`

	/*
	   The window to perform maintenance in.
	   Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00". See [RDS
	   Maintenance Window
	   docs](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow)
	   for more information.
	*/
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	// Amount of time in days to retain Performance Insights data. Valid values are `7`, `731` (2 years) or a multiple of `31`. When specifying `performance_insights_retention_period`, `performance_insights_enabled` needs to be set to true. Defaults to '7'.
	PerformanceInsightsRetentionPeriod int `json:"performanceInsightsRetentionPeriod,omitempty" yaml:"performanceInsightsRetentionPeriod,omitempty"`

	/*
	   Specifies whether the DB instance is
	   encrypted. Note that if you are creating a cross-region read replica this field
	   is ignored and you should instead declare `kms_key_id` with a valid ARN. The
	   default is `false` if not specified.
	*/
	StorageEncrypted bool `json:"storageEncrypted,omitempty" yaml:"storageEncrypted,omitempty"`

	// The identifier of the CA certificate for the DB instance.
	CaCertIdentifier string `json:"caCertIdentifier,omitempty" yaml:"caCertIdentifier,omitempty"`

	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix string `json:"identifierPrefix,omitempty" yaml:"identifierPrefix,omitempty"`

	/*
	   The ARN for the KMS encryption key. If creating an
	   encrypted replica, set this to the destination KMS ARN.
	*/
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	/*
	   Indicates that major version
	   upgrades are allowed. Changing this parameter does not result in an outage and
	   the change is asynchronously applied as soon as possible.
	*/
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" yaml:"allowMajorVersionUpgrade,omitempty"`

	// The name of the RDS instance, if omitted, this provider will assign a random, unique identifier. Required if `restore_to_point_in_time` is specified.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// Name of the DB option group to associate.
	OptionGroupName string `json:"optionGroupName,omitempty" yaml:"optionGroupName,omitempty"`

	// The ARN for the KMS key to encrypt Performance Insights data. When specifying `performance_insights_kms_key_id`, `performance_insights_enabled` needs to be set to true. Once KMS key is set, it can never be changed.
	PerformanceInsightsKmsKeyId string `json:"performanceInsightsKmsKeyId,omitempty" yaml:"performanceInsightsKmsKeyId,omitempty"`

	/*
	   Bool to control if instance is publicly
	   accessible. Default is `false`.
	*/
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	// If the DB instance should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// The ARN for the Secrets Manager secret with the self managed Active Directory credentials for the user joining the domain. Conflicts with `domain` and `domain_iam_role_name`.
	DomainAuthSecretArn string `json:"domainAuthSecretArn,omitempty" yaml:"domainAuthSecretArn,omitempty"`

	/*
	   The amount of provisioned IOPS. Setting this implies a
	   storage_type of "io1". Can only be set when `storage_type` is `"io1"` or `"gp3"`.
	   Cannot be specified for gp3 storage if the `allocated_storage` value is below a per-`engine` threshold.
	   See the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#gp3-storage) for details.
	*/
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// The storage throughput value for the DB instance. Can only be set when `storage_type` is `"gp3"`. Cannot be specified if the `allocated_storage` value is below a per-`engine` threshold. See the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#gp3-storage) for details.
	StorageThroughput int `json:"storageThroughput,omitempty" yaml:"storageThroughput,omitempty"`

	/*
	   Specifies whether mappings of AWS Identity and Access Management (IAM) accounts to database
	   accounts is enabled.
	*/
	IamDatabaseAuthenticationEnabled bool `json:"iamDatabaseAuthenticationEnabled,omitempty" yaml:"iamDatabaseAuthenticationEnabled,omitempty"`

	// The instance type of the RDS instance.
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`

	// Specifies whether Performance Insights are enabled. Defaults to false.
	PerformanceInsightsEnabled bool `json:"performanceInsightsEnabled,omitempty" yaml:"performanceInsightsEnabled,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies where automated backups and manual snapshots are stored. Possible values are `region` (default) and `outposts`. See [Working with Amazon RDS on AWS Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) for more information.
	BackupTarget string `json:"backupTarget,omitempty" yaml:"backupTarget,omitempty"`

	/*
	   The character set name to use for DB encoding in Oracle and Microsoft SQL instances (collation).
	   This can't be changed.
	   See [Oracle Character Sets Supported in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html) or
	   [Server-Level Collation for Microsoft SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Collation.html) for more information.
	   Cannot be set  with `replicate_source_db`, `restore_to_point_in_time`, `s3_import`, or `snapshot_identifier`.
	*/
	CharacterSetName string `json:"characterSetName,omitempty" yaml:"characterSetName,omitempty"`

	// The IPv4 DNS IP addresses of your primary and secondary self managed Active Directory domain controllers. Two IP addresses must be provided. If there isn't a secondary domain controller, use the IP address of the primary domain controller for both entries in the list. Conflicts with `domain` and `domain_iam_role_name`.
	DomainDnsIps []string `json:"domainDnsIps,omitempty" yaml:"domainDnsIps,omitempty"`

	// The self managed Active Directory organizational unit for your DB instance to join. Conflicts with `domain` and `domain_iam_role_name`.
	DomainOu string `json:"domainOu,omitempty" yaml:"domainOu,omitempty"`

	// The engine version to use. If `auto_minor_version_upgrade` is enabled, you can provide a prefix of the version such as `8.0` (for `8.0.36`). The actual engine version used is returned in the attribute `engine_version_actual`, see Attribute Reference below. For supported values, see the EngineVersion parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html). Note that for Amazon Aurora instances the engine version must match the DB cluster's engine version'.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Set to true to allow RDS to manage the master user password in Secrets Manager. Cannot be set if `password` is provided.
	ManageMasterUserPassword bool `json:"manageMasterUserPassword,omitempty" yaml:"manageMasterUserPassword,omitempty"`

	/*
	   Indicates that minor engine upgrades
	   will be applied automatically to the DB instance during the maintenance window.
	   Defaults to true.
	*/
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	/*
	   Enables low-downtime updates using [RDS Blue/Green deployments][blue-green].
	   See `blue_green_update` below.
	*/
	BlueGreenUpdate types.Rds_InstanceBlueGreenUpdate `json:"blueGreenUpdate,omitempty" yaml:"blueGreenUpdate,omitempty"`

	// The life cycle type for this DB instance. This setting applies only to RDS for MySQL and RDS for PostgreSQL. Valid values are `open-source-rds-extended-support`, `open-source-rds-extended-support-disabled`. Default value is `open-source-rds-extended-support`. [Using Amazon RDS Extended Support]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html
	EngineLifecycleSupport string `json:"engineLifecycleSupport,omitempty" yaml:"engineLifecycleSupport,omitempty"`

	/*
	   Specifies whether the replica is in either `mounted` or `open-read-only` mode. This attribute
	   is only supported by Oracle instances. Oracle replicas operate in `open-read-only` mode unless otherwise specified. See [Working with Oracle Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) for more information.
	*/
	ReplicaMode string `json:"replicaMode,omitempty" yaml:"replicaMode,omitempty"`

	/*
	   One of "standard" (magnetic), "gp2" (general
	   purpose SSD), "gp3" (general purpose SSD that needs `iops` independently)
	   or "io1" (provisioned IOPS SSD). The default is "io1" if `iops` is specified,
	   "gp2" if not.
	*/
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	/*
	   Specifies whether any database modifications
	   are applied immediately, or during the next maintenance window. Default is
	   `false`. See [Amazon RDS Documentation for more
	   information.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
	*/
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// Copy all Instance `tags` to snapshots. Default is `false`.
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" yaml:"copyTagsToSnapshot,omitempty"`

	/*
	   The interval, in seconds, between points
	   when Enhanced Monitoring metrics are collected for the DB instance. To disable
	   collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid
	   Values: 0, 1, 5, 10, 15, 30, 60.
	*/
	MonitoringInterval int `json:"monitoringInterval,omitempty" yaml:"monitoringInterval,omitempty"`

	// Specifies whether to remove automated backups immediately after the DB instance is deleted. Default is `true`.
	DeleteAutomatedBackups bool `json:"deleteAutomatedBackups,omitempty" yaml:"deleteAutomatedBackups,omitempty"`

	/*
	   License model information for this DB instance. Valid values for this field are as follows:
	   - RDS for MariaDB: `general-public-license`
	   - RDS for Microsoft SQL Server: `license-included`
	   - RDS for MySQL: `general-public-license`
	   - RDS for Oracle: `bring-your-own-license | license-included`
	   - RDS for PostgreSQL: `postgresql-license`
	*/
	LicenseModel string `json:"licenseModel,omitempty" yaml:"licenseModel,omitempty"`

	// The fully qualified domain name (FQDN) of the self managed Active Directory domain. Conflicts with `domain` and `domain_iam_role_name`.
	DomainFqdn string `json:"domainFqdn,omitempty" yaml:"domainFqdn,omitempty"`

	// The database engine to use. For supported values, see the Engine parameter in [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html). Note that for Amazon Aurora instances the engine must match the DB cluster's engine'. For information on the difference between the available Aurora MySQL engines see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html) in the Amazon RDS User Guide.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Specifies if the RDS instance is multi-AZ
	MultiAz bool `json:"multiAz,omitempty" yaml:"multiAz,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A configuration block for restoring a DB instance to an arbitrary point in time. Requires the `identifier` argument to be set with the name of the new DB instance to be created. See Restore To Point In Time below for details.
	RestoreToPointInTime types.Rds_InstanceRestoreToPointInTime `json:"restoreToPointInTime,omitempty" yaml:"restoreToPointInTime,omitempty"`

	/*
	   Specifies whether or not to create this
	   database from a snapshot. This correlates to the snapshot ID you'd find in the
	   RDS console, e.g: rds:production-2015-06-26-06-05.
	*/
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" yaml:"snapshotIdentifier,omitempty"`

	// The AZ for the RDS instance.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	/*
	   Name of DB subnet group. DB instance will
	   be created in the VPC associated with the DB subnet group. If unspecified, will
	   be created in the `default` VPC, or in EC2 Classic, if available. When working
	   with read replicas, it should be specified only if the source database
	   specifies an instance in another AWS Region. See [DBSubnetGroupName in API
	   action CreateDBInstanceReadReplica](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstanceReadReplica.html)
	   for additional read replica constraints.
	*/
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" yaml:"dbSubnetGroupName,omitempty"`

	/*
	   Time zone of the DB instance. `timezone` is currently
	   only supported by Microsoft SQL Server. The `timezone` can only be set on
	   creation. See [MSSQL User
	   Guide](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone)
	   for more information.
	*/
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`

	/*
	   The daily time range (in UTC) during which automated backups are created if they are enabled.
	   Example: "09:46-10:16". Must not overlap with `maintenance_window`.
	*/
	BackupWindow string `json:"backupWindow,omitempty" yaml:"backupWindow,omitempty"`

	// The ID of the Directory Service Active Directory domain to create the instance in. Conflicts with `domain_fqdn`, `domain_ou`, `domain_auth_secret_arn` and a `domain_dns_ips`.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
