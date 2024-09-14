package types

type Elasticache_ClusterCacheNode struct {
	// The port number on which each of the cache nodes will accept connections. For Memcached the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`. Changing this value will re-create the resource.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	//
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	//
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`
}
