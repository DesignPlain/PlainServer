package elasticache

type GlobalReplicationGroup struct {
	// The number of node groups (shards) on the global replication group.
	NumNodeGroups int `json:"numNodeGroups,omitempty" yaml:"numNodeGroups,omitempty"`

	/*
	   An ElastiCache Parameter Group to use for the Global Replication Group.
	   Required when upgrading a major engine version, but will be ignored if left configured after the upgrade is complete.
	   Specifying without a major version upgrade will fail.
	   Note that ElastiCache creates a copy of this parameter group for each member replication group.
	*/
	ParameterGroupName string `json:"parameterGroupName,omitempty" yaml:"parameterGroupName,omitempty"`

	// The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primary_replication_group_id` is changed, creates a new resource.
	PrimaryReplicationGroupId string `json:"primaryReplicationGroupId,omitempty" yaml:"primaryReplicationGroupId,omitempty"`

	/*
	   Specifies whether read-only replicas will be automatically promoted to read/write primary if the existing primary fails.
	   When creating, by default the Global Replication Group inherits the automatic failover setting of the primary replication group.
	*/
	AutomaticFailoverEnabled bool `json:"automaticFailoverEnabled,omitempty" yaml:"automaticFailoverEnabled,omitempty"`

	/*
	   The instance class used.
	   See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html)
	   and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
	   When creating, by default the Global Replication Group inherits the node type of the primary replication group.
	*/
	CacheNodeType string `json:"cacheNodeType,omitempty" yaml:"cacheNodeType,omitempty"`

	/*
	   Redis version to use for the Global Replication Group.
	   When creating, by default the Global Replication Group inherits the version of the primary replication group.
	   If a version is specified, the Global Replication Group and all member replication groups will be upgraded to this version.
	   Cannot be downgraded without replacing the Global Replication Group and all member replication groups.
	   When the version is 7 or higher, the major and minor version should be set, e.g., `7.2`.
	   When the version is 6, the major and minor version can be set, e.g., `6.2`,
	   or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
	   The actual engine version used is returned in the attribute `engine_version_actual`, see Attribute Reference below.
	*/
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// A user-created description for the global replication group.
	GlobalReplicationGroupDescription string `json:"globalReplicationGroupDescription,omitempty" yaml:"globalReplicationGroupDescription,omitempty"`

	// The suffix name of a Global Datastore. If `global_replication_group_id_suffix` is changed, creates a new resource.
	GlobalReplicationGroupIdSuffix string `json:"globalReplicationGroupIdSuffix,omitempty" yaml:"globalReplicationGroupIdSuffix,omitempty"`
}
