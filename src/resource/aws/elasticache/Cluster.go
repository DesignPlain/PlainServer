package elasticache

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	// Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`.
	AzMode string `json:"azMode,omitempty" yaml:"azMode,omitempty"`

	/*
	   Specifies the weekly time range for when maintenance
	   on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	   The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`.
	*/
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	// List of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `num_cache_nodes`. If you want all the nodes in the same Availability Zone, use `availability_zone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones []string `json:"preferredAvailabilityZones,omitempty" yaml:"preferredAvailabilityZones,omitempty"`

	// Group identifier. ElastiCache converts this name to lowercase. Changing this value will re-create the resource.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	/*
	   The name of the parameter group to associate with this cache cluster.

	   The following arguments are optional:
	*/
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	/*
	   Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
	   Only supported for engine type `"redis"` and if the engine version is 6 or higher.
	   Defaults to `true`.
	*/
	AutoMinorVersionUpgrade string `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// Specifies the destination and format of Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See the documentation on [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html). See Log Delivery Configuration below for more details.
	LogDeliveryConfigurations []types.Elasticache_ClusterLogDeliveryConfiguration `json:"logDeliveryConfigurations,omitempty" yaml:"logDeliveryConfigurations,omitempty"`

	// Name of a snapshot from which to restore data into the new node group. Changing `snapshot_name` forces a new resource.
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`

	// Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off. Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro cache nodes
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" yaml:"snapshotRetentionLimit,omitempty"`

	// Whether any database modifications are applied immediately, or during the next maintenance window. Default is `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html).
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	/*
	   Version number of the cache engine to be used.
	   If not set, defaults to the latest version.
	   See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html) in the AWS Documentation for supported versions.
	   When `engine` is `redis` and the version is 7 or higher, the major and minor version should be set, e.g., `7.2`.
	   When the version is 6, the major and minor version can be set, e.g., `6.2`,
	   or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
	   Otherwise, specify the full version desired, e.g., `5.0.6`.
	   The actual engine version used is returned in the attribute `engine_version_actual`, see Attribute Reference below. Cannot be provided with `replication_group_id.`
	*/
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// The instance class used. See AWS documentation for information on [supported node types for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html). See AWS documentation for information on [supported node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/nodes-select-size.html). For Memcached, changing this value will re-create the resource.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	// The port number on which each of the cache nodes will accept connections. For Memcached the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`. Changing this value will re-create the resource.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of your final cluster snapshot. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// The IP versions for cache cluster connections. IPv6 is supported with Redis engine `6.2` onword or Memcached version `1.6.6` for all [Nitro system](https://aws.amazon.com/ec2/nitro/) instances. Valid values are `ipv4`, `ipv6` or `dual_stack`.
	NetworkType string `json:"networkType,omitempty" yaml:"networkType,omitempty"`

	// The initial number of cache nodes that the cache cluster will have. For Redis, this value must be 1. For Memcached, this value must be between 1 and 40. If this number is reduced on subsequent runs, the highest numbered nodes will be removed.
	NumCacheNodes int `json:"numCacheNodes,omitempty" yaml:"numCacheNodes,omitempty"`

	// One or more VPC security groups associated with the cache cluster. Cannot be provided with `replication_group_id.`
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Name of the subnet group to be used for the cache cluster. Changing this value will re-create the resource. Cannot be provided with `replication_group_id.`
	SubnetGroupName string `json:"subnetGroupName,omitempty" yaml:"subnetGroupName,omitempty"`

	// Enable encryption in-transit. Supported only with Memcached versions `1.6.12` and later, running in a VPC. See the [ElastiCache in-transit encryption](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/in-transit-encryption-mc.html) documentation for more details.
	TransitEncryptionEnabled bool `json:"transitEncryptionEnabled,omitempty" yaml:"transitEncryptionEnabled,omitempty"`

	// Name of the cache engine to be used for this cache cluster. Valid values are `memcached` or `redis`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// The IP version to advertise in the discovery protocol. Valid values are `ipv4` or `ipv6`.
	IpDiscovery string `json:"ipDiscovery,omitempty" yaml:"ipDiscovery,omitempty"`

	// ARN of an SNS topic to send ElastiCache notifications to. Example: `arn:aws:sns:us-east-1:012345678999:my_sns_topic`.
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" yaml:"notificationTopicArn,omitempty"`

	// Single-element string list containing an Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3. The object name cannot contain any commas. Changing `snapshot_arns` forces a new resource.
	SnapshotArns string `json:"snapshotArns,omitempty" yaml:"snapshotArns,omitempty"`

	// Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow string `json:"snapshotWindow,omitempty" yaml:"snapshotWindow,omitempty"`

	// Specify the outpost mode that will apply to the cache cluster creation. Valid values are `"single-outpost"` and `"cross-outpost"`, however AWS currently only supports `"single-outpost"` mode.
	OutpostMode string `json:"outpostMode,omitempty" yaml:"outpostMode,omitempty"`

	// The outpost ARN in which the cache cluster will be created.
	PreferredOutpostArn string `json:"preferredOutpostArn,omitempty" yaml:"preferredOutpostArn,omitempty"`

	// ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId string `json:"replicationGroupId,omitempty" yaml:"replicationGroupId,omitempty"`
}
