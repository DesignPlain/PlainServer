package types

type Memorydb_ClusterShardNodeEndpoint struct {
	// DNS hostname of the node.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The port number on which each of the nodes accepts connections. Defaults to `6379`.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
