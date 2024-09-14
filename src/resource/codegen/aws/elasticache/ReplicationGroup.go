package elasticache

import types "libds/aws/types"

type ReplicationGroup struct {
	/*
	   Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
	   Only supported for engine type `"redis"` and if the engine version is 6 or higher.
	   Defaults to `true`.
	*/
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" yaml:"autoMinorVersionUpgrade,omitempty"`

	// The IP version to advertise in the discovery protocol. Valid values are `ipv4` or `ipv6`.
	IpDiscovery string `json:"ipDiscovery,omitempty" yaml:"ipDiscovery,omitempty"`

	// The IP versions for cache cluster connections. Valid values are `ipv4`, `ipv6` or `dual_stack`.
	NetworkType string `json:"networkType,omitempty" yaml:"networkType,omitempty"`

	// Name of the cache subnet group to be used for the replication group.
	SubnetGroupName string `json:"subnetGroupName,omitempty" yaml:"subnetGroupName,omitempty"`

	// Map of tags to assign to the resource. Adding tags to this resource will add or overwrite any existing tags on the clusters in the replication group and not to the group itself. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Whether to enable encryption in transit.
	   Changing this argument with an `engine_version` < `7.0.5` will force a replacement.
	   Engine versions prior to `7.0.5` only allow this transit encryption to be configured during creation of the replication group.
	*/
	TransitEncryptionEnabled bool `json:"transitEncryptionEnabled,omitempty" yaml:"transitEncryptionEnabled,omitempty"`

	// Whether to enable encryption at rest.
	AtRestEncryptionEnabled bool `json:"atRestEncryptionEnabled,omitempty" yaml:"atRestEncryptionEnabled,omitempty"`

	// Password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
	AuthToken string `json:"authToken,omitempty" yaml:"authToken,omitempty"`

	// Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to `true` when using r6gd nodes.
	DataTieringEnabled bool `json:"dataTieringEnabled,omitempty" yaml:"dataTieringEnabled,omitempty"`

	// Specifies the weekly time range for when maintenance on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	/*
	   Specifies whether to enable Multi-AZ Support for the replication group.
	   If `true`, `automatic_failover_enabled` must also be enabled.
	   Defaults to `false`.
	*/
	MultiAzEnabled bool `json:"multiAzEnabled,omitempty" yaml:"multiAzEnabled,omitempty"`

	/*
	   Number of cache clusters (primary and replicas) this replication group will have.
	   If `automatic_failover_enabled` or `multi_az_enabled` are `true`, must be at least 2.
	   Updates will occur before other modifications.
	   Conflicts with `num_node_groups` and `replicas_per_node_group`.
	   Defaults to `1`.
	*/
	NumCacheClusters int `json:"numCacheClusters,omitempty" yaml:"numCacheClusters,omitempty"`

	// User Group ID to associate with the replication group. Only a maximum of one (1) user group ID is valid. --NOTE:-- This argument _is_ a set because the AWS specification allows for multiple IDs. However, in practice, AWS only allows a maximum size of one.
	UserGroupIds []string `json:"userGroupIds,omitempty" yaml:"userGroupIds,omitempty"`

	// IDs of one or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If enabled, `num_cache_clusters` must be greater than 1. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
	AutomaticFailoverEnabled bool `json:"automaticFailoverEnabled,omitempty" yaml:"automaticFailoverEnabled,omitempty"`

	// User-created description for the replication group. Must not be empty.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of your final node group (shard) snapshot. ElastiCache creates the snapshot from the primary node in the cluster. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" yaml:"finalSnapshotIdentifier,omitempty"`

	// The ID of the global replication group to which this replication group should belong. If this parameter is specified, the replication group is added to the specified global replication group as a secondary replication group; otherwise, the replication group is not part of any global replication group. If `global_replication_group_id` is set, the `num_node_groups` parameter cannot be set.
	GlobalReplicationGroupId string `json:"globalReplicationGroupId,omitempty" yaml:"globalReplicationGroupId,omitempty"`

	// ARN of an SNS topic to send ElastiCache notifications to. Example: `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" yaml:"notificationTopicArn,omitempty"`

	/*
	   Number of node groups (shards) for this Redis replication group.
	   Changing this number will trigger a resizing operation before other settings modifications.
	   Conflicts with `num_cache_clusters`.
	*/
	NumNodeGroups int `json:"numNodeGroups,omitempty" yaml:"numNodeGroups,omitempty"`

	// Port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Specifies the destination and format of Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See the documentation on [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See Log Delivery Configuration below for more details.
	LogDeliveryConfigurations []types.Elasticache_ReplicationGroupLogDeliveryConfiguration `json:"logDeliveryConfigurations,omitempty" yaml:"logDeliveryConfigurations,omitempty"`

	// List of ARNs that identify Redis RDB snapshot files stored in Amazon S3. The names object names cannot contain any commas.
	SnapshotArns []string `json:"snapshotArns,omitempty" yaml:"snapshotArns,omitempty"`

	// Name of the cache engine to be used for the clusters in this replication group. The only valid value is `redis`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `at_rest_encryption_enabled = true`.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	/*
	   Version number of the cache engine to be used for the cache clusters in this replication group.
	   If the version is 7 or higher, the major and minor version should be set, e.g., `7.2`.
	   If the version is 6, the major and minor version can be set, e.g., `6.2`,
	   or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
	   Otherwise, specify the full version desired, e.g., `5.0.6`.
	   The actual engine version used is returned in the attribute `engine_version_actual`, see Attribute Reference below.
	*/
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used. To enable "cluster mode", i.e., data sharding, use a parameter group that has the parameter `cluster-enabled` set to true.
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	// Names of one or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud.
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" yaml:"securityGroupNames,omitempty"`

	// Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of `snapshot_retention_limit` is set to zero (0), backups are turned off. Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro cache nodes
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" yaml:"snapshotRetentionLimit,omitempty"`

	/*
	   A setting that enables clients to migrate to in-transit encryption with no downtime.
	   Valid values are `preferred` and `required`.
	   When enabling encryption on an existing replication group, this must first be set to `preferred` before setting it to `required` in a subsequent apply.
	   See the `TransitEncryptionMode` field in the [`CreateReplicationGroup` API documentation](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateReplicationGroup.html) for additional details.
	*/
	TransitEncryptionMode string `json:"transitEncryptionMode,omitempty" yaml:"transitEncryptionMode,omitempty"`

	// Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately bool `json:"applyImmediately,omitempty" yaml:"applyImmediately,omitempty"`

	// Strategy to use when updating the `auth_token`. Valid values are `SET`, `ROTATE`, and `DELETE`. Defaults to `ROTATE`.
	AuthTokenUpdateStrategy string `json:"authTokenUpdateStrategy,omitempty" yaml:"authTokenUpdateStrategy,omitempty"`

	// Specifies whether cluster mode is enabled or disabled. Valid values are `enabled` or `disabled` or `compatible`
	ClusterMode string `json:"clusterMode,omitempty" yaml:"clusterMode,omitempty"`

	// Instance class to be used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html). Required unless `global_replication_group_id` is set. Cannot be set if `global_replication_group_id` is set.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	// List of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is considered. The first item in the list will be the primary node. Ignored when updating.
	PreferredCacheClusterAzs []string `json:"preferredCacheClusterAzs,omitempty" yaml:"preferredCacheClusterAzs,omitempty"`

	// Name of a snapshot from which to restore data into the new node group. Changing the `snapshot_name` forces a new resource.
	SnapshotName string `json:"snapshotName,omitempty" yaml:"snapshotName,omitempty"`

	/*
	   Number of replica nodes in each node group.
	   Changing this number will trigger a resizing operation before other settings modifications.
	   Valid values are 0 to 5.
	   Conflicts with `num_cache_clusters`.
	   Can only be set if `num_node_groups` is set.
	*/
	ReplicasPerNodeGroup int `json:"replicasPerNodeGroup,omitempty" yaml:"replicasPerNodeGroup,omitempty"`

	/*
	   Replication group identifier. This parameter is stored as a lowercase string.

	   The following arguments are optional:
	*/
	ReplicationGroupId string `json:"replicationGroupId,omitempty" yaml:"replicationGroupId,omitempty"`

	// Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
	SnapshotWindow string `json:"snapshotWindow,omitempty" yaml:"snapshotWindow,omitempty"`
}
