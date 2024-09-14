package types

type Memorydb_getClusterShardNode struct {
	// The Availability Zone in which the node resides.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The date and time when the node was created. Example: `2022-01-01T21:00:00Z`.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	//
	Endpoints []Memorydb_getClusterShardNodeEndpoint `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`

	// Name of the cluster.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
