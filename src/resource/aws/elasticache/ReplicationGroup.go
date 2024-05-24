package elasticache

import types "DesignSphere_Server/src/resource/aws/types"

type ReplicationGroup struct {
	// Names of one or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud.
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" yaml:"securityGroupNames,omitempty"`

	// List of ARNs that identify Redis RDB snapshot files stored in Amazon S3. The names object names cannot contain any commas.
	SnapshotArns []string `json:"snapshotArns,omitempty" yaml:"snapshotArns,omitempty"`

	// Strategy to use when updating the `auth_token`. Valid values are `SET`, `ROTATE`, and `DELETE`. Defaults to `ROTATE`.
	AuthTokenUpdateStrategy string `json:"authTokenUpdateStrategy,omitempty" yaml:"authTokenUpdateStrategy,omitempty"`

	// User-created description for the replication group. Must not be empty.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Replication group identifier. This parameter is stored as a lowercase string.

	   The following arguments are optional:
	*/
	ReplicationGroupId string `json:"replicationGroupId,omitempty" yaml:"replicationGroupId,omitempty"`

	/*
	   Number of replica nodes in each node group.
	   Changing this number will trigger a resizing operation before other settings modifications.
	   Valid values are 0 to 5.
	*/
	ReplicasPerNodeGroup int `json:"replicasPerNodeGroup,omitempty" yaml:"replicasPerNodeGroup,omitempty"`

	// Map of tags to assign to the resource. Adding tags to this resource will add or overwrite any existing tags on the clusters in the replication group and not to the group itself. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to `true` when using r6gd nodes.
	DataTieringEnabled bool `json:"dataTieringEnabled,omitempty" yaml:"dataTieringEnabled,omitempty"`

	// ARN of an SNS topic to send ElastiCache notifications to. Example: `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" yaml:"notificationTopicArn,omitempty"`

	// List of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is considered. The first item in the list will be the primary node. Ignored when updating.
	PreferredCacheClusterAzs []string `json:"preferredCacheClusterAzs,omitempty" yaml:"preferredCacheClusterAzs,omitempty"`

	// Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
	SnapshotWindow string `json:"snapshotWindow,omitempty" yaml:"snapshotWindow,omitempty"`

	// Whether to enable encryption in transit.
	TransitEncryptionEnabled bool `json:"transitEncryptionEnabled,omitempty" yaml:"transitEncryptionEnabled,omitempty"`

	// User Group ID to associate with the replication group. Only a maximum of one (1) user group ID is valid. --NOTE:-- This argument _is_ a set because the AWS specification allows for multiple IDs. However, in practice, AWS only allows a maximum size of one.
	UserGroupIds []string `json:"userGroupIds,omitempty" yaml:"userGroupIds,omitempty"`

	// Specifies the destination and format of Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See the documentation on [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See Log Delivery Configuration below for more details.
	LogDeliveryConfigurations []types.Elasticache_ReplicationGroupLogDeliveryConfiguration `json:"logDeliveryConfigurations,omitempty" yaml:"logDeliveryConfigurations,omitempty"`

	// IDs of one or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Name of the cache subnet group to be used for the replication group.
	SubnetGroupName string `json:"subnetGroupName,omitempty" yaml:"subnetGroupName,omitempty"`

	/*
	   Number of node groups (shards) for this Redis replication group.
	   Changing this number will trigger a resizing operation before other settings modifications.
	*/
	NumNodeGroups int `json:"numNodeGroups,omitempty" yaml:"numNodeGroups,omitempty"`

	// Whether to enable encryption at rest.
	AtRestEncryptionEnabled bool `json:"atRestEncryptionEnabled,omitempty" yaml:"atRestEncryptionEnabled,omitempty"`

	// Name of the cache engine to be used for the clusters in this replication group. The only valid value is `redis`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Specifies the weekly time range for when maintenance on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	// Name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used. To enable "cluster mode", i.e., data sharding, use a parameter group that has the parameter `cluster-enabled` set to true.
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	// Port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of `snapshot_retention_limit` is set to zero (0), backups are turned off. Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro cache nodes
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" yaml:"snapshotRetentionLimit,omitempty"`

	// Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	/*
	   Version number of the cache engine to be used for the cache clusters in this replication group.
	   If the version is 7 or higher, the major and minor version should be set, e.g., `7.2`.
	   If the version is 6, the major and minor version can be set, e.g., `6.2`,
	   or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
	   Otherwise, specify the full version desired, e.g., `5.0.6`.
	   The actual engine version used is returned in the attribute `engine_version_actual`, see Attribute Reference below.
	*/
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// The ID of the global replication group to which this replication group should belong. If this parameter is specified, the replication group is added to the specified global replication group as a secondary replication group; otherwise, the replication group is not part of any global replication group. If `global_replication_group_id` is set, the `num_node_groups` parameter cannot be set.
	GlobalReplicationGroupId string `json:"globalReplicationGroupId,omitempty" yaml:"globalReplicationGroupId,omitempty"`

	// The name of your final node group (shard) snapshot. ElastiCache creates the snapshot from the primary node in the cluster. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// The IP version to advertise in the discovery protocol. Valid values are `ipv4` or `ipv6`.
	IpDiscovery string `json:"ipDiscovery,omitempty" yaml:"ipDiscovery,omitempty"`

	// The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `at_rest_encryption_enabled = true`.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Specifies whether to enable Multi-AZ Support for the replication group. If `true`, `automatic_failover_enabled` must also be enabled. Defaults to `false`.
	MultiAzEnabled bool `json:"multiAzEnabled,omitempty" yaml:"multiAzEnabled,omitempty"`

	// The IP versions for cache cluster connections. Valid values are `ipv4`, `ipv6` or `dual_stack`.
	NetworkType string `json:"networkType,omitempty" yaml:"networkType,omitempty"`

	// Password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
	AuthToken string `json:"authToken,omitempty" yaml:"authToken,omitempty"`

	/*
	   Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
	   Only supported for engine type `"redis"` and if the engine version is 6 or higher.
	   Defaults to `true`.
	*/
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If enabled, `num_cache_clusters` must be greater than 1. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
	AutomaticFailoverEnabled bool `json:"automaticFailoverEnabled,omitempty" yaml:"automaticFailoverEnabled,omitempty"`

	// Instance class to be used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html). Required unless `global_replication_group_id` is set. Cannot be set if `global_replication_group_id` is set.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	// Number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications. Conflicts with `num_node_groups`. Defaults to `1`.
	NumCacheClusters int `json:"numCacheClusters,omitempty" yaml:"numCacheClusters,omitempty"`

	// Name of a snapshot from which to restore data into the new node group. Changing the `snapshot_name` forces a new resource.
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`
}
