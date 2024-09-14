package elasticache

import types "libds/aws/types"

type ServerlessCache struct {
	// Sets the cache usage limits for storage and ElastiCache Processing Units for the cache. See configuration below.
	CacheUsageLimits types.Elasticache_ServerlessCacheCacheUsageLimits `json:"cacheUsageLimits,omitempty" yaml:"cacheUsageLimits,omitempty"`

	// ARN of the customer managed key for encrypting the data at rest. If no KMS key is provided, a default service key is used.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	/*
	   The version of the cache engine that will be used to create the serverless cache.
	   See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html) in the AWS Documentation for supported versions.
	*/
	MajorEngineVersion string `json:"majorEngineVersion,omitempty" yaml:"majorEngineVersion,omitempty"`

	// A list of the one or more VPC security groups to be associated with the serverless cache. The security group will authorize traffic access for the VPC end-point (private-link). If no other information is given this will be the VPC’s Default Security Group that is associated with the cluster VPC end-point.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The list of ARN(s) of the snapshot that the new serverless cache will be created from. Available for Redis only.
	SnapshotArnsToRestores []string `json:"snapshotArnsToRestores,omitempty" yaml:"snapshotArnsToRestores,omitempty"`

	//
	Timeouts types.Elasticache_ServerlessCacheTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The daily time that snapshots will be created from the new serverless cache. Only supported for engine type `"redis"`. Defaults to `0`.
	DailySnapshotTime string `json:"dailySnapshotTime,omitempty" yaml:"dailySnapshotTime,omitempty"`

	// User-provided description for the serverless cache. The default is NULL.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The Cluster name which serves as a unique identifier to the serverless cache

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The number of snapshots that will be retained for the serverless cache that is being created. As new snapshots beyond this limit are added, the oldest snapshots will be deleted on a rolling basis. Available for Redis only.
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" yaml:"snapshotRetentionLimit,omitempty"`

	// The identifier of the UserGroup to be associated with the serverless cache. Available for Redis only. Default is NULL.
	UserGroupId string `json:"userGroupId,omitempty" yaml:"userGroupId,omitempty"`

	// Name of the cache engine to be used for this cache cluster. Valid values are `memcached` or `redis`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// A list of the identifiers of the subnets where the VPC endpoint for the serverless cache will be deployed. All the subnetIds must belong to the same VPC.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
