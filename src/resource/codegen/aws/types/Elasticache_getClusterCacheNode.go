package types

type Elasticache_getClusterCacheNode struct {
	//
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Availability Zone for the cache cluster.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	//
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	//
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`

	/*
	   The port number on which each of the cache nodes will
	   accept connections.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
