package rds

type ClusterInstance struct {
	// ARN for the KMS key to encrypt Performance Insights data. When specifying `performance_insights_kms_key_id`, `performance_insights_enabled` needs to be set to true.
	PerformanceInsightsKmsKeyId string `json:"performanceInsightsKmsKeyId,omitempty" yaml:"performanceInsightsKmsKeyId,omitempty"`

	// Amount of time in days to retain Performance Insights data. Valid values are `7`, `731` (2 years) or a multiple of `31`. When specifying `performance_insights_retention_period`, `performance_insights_enabled` needs to be set to true. Defaults to '7'.
	PerformanceInsightsRetentionPeriod int `json:"performanceInsightsRetentionPeriod,omitempty" yaml:"performanceInsightsRetentionPeriod,omitempty"`

	// Identifier of the `aws.rds.Cluster` in which to launch this instance.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// DB subnet group to associate with this DB instance. --NOTE:-- This must match the `db_subnet_group_name` of the attached `aws.rds.Cluster`.
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" yaml:"dbSubnetGroupName,omitempty"`

	// Interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval int `json:"monitoringInterval,omitempty" yaml:"monitoringInterval,omitempty"`

	// Identifier for the RDS instance, if omitted, Pulumi will assign a random, unique identifier.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// Instance class to use. For details on CPU and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.-` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`

	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled bool `json:"performanceInsightsEnabled,omitempty" yaml:"performanceInsightsEnabled,omitempty"`

	// Daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00". --NOTE:-- If `preferred_backup_window` is set at the cluster level, this argument --must-- be omitted.
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" yaml:"preferredBackupWindow,omitempty"`

	// Window to perform maintenance in. Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" yaml:"preferredMaintenanceWindow,omitempty"`

	// Specifies whether any database modifications are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// Identifier of the CA certificate for the DB instance.
	CaCertIdentifier string `json:"caCertIdentifier,omitempty" yaml:"caCertIdentifier,omitempty"`

	/*
	   Name of the database engine to be used for the RDS cluster instance.
	   Valid Values: `aurora-mysql`, `aurora-postgresql`.
	*/
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix string `json:"identifierPrefix,omitempty" yaml:"identifierPrefix,omitempty"`

	// Bool to control if instance is publicly accessible. Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more details on controlling this property.
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" yaml:"publiclyAccessible,omitempty"`

	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" yaml:"copyTagsToSnapshot,omitempty"`

	// Instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	CustomIamInstanceProfile string `json:"customIamInstanceProfile,omitempty" yaml:"customIamInstanceProfile,omitempty"`

	// Name of the DB parameter group to associate with this instance.
	DbParameterGroupName string `json:"dbParameterGroupName,omitempty" yaml:"dbParameterGroupName,omitempty"`

	// ARN for the IAM role that permits RDS to send enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html) what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	MonitoringRoleArn string `json:"monitoringRoleArn,omitempty" yaml:"monitoringRoleArn,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier int `json:"promotionTier,omitempty" yaml:"promotionTier,omitempty"`

	// Map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Database engine version.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`
}
