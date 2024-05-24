package types

type Memorydb_ClusterShard struct {
	// Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of nodes in this shard.
	Nodes []Memorydb_ClusterShardNode `json:"nodes,omitempty" yaml:"nodes,omitempty"`

	// Number of individual nodes in this shard.
	NumNodes int `json:"numNodes,omitempty" yaml:"numNodes,omitempty"`

	// Keyspace for this shard. Example: `0-16383`.
	Slots string `json:"slots,omitempty" yaml:"slots,omitempty"`
}
