package types

type Memorydb_ClusterShardNode struct {
	// The Availability Zone in which the node resides.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The date and time when the node was created. Example: `2022-01-01T21:00:00Z`.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	//
	Endpoints []Memorydb_ClusterShardNodeEndpoint `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`

	// Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
