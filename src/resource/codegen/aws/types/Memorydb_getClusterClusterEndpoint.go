package types

type Memorydb_getClusterClusterEndpoint struct {
	// DNS hostname of the node.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Port number that this node is listening on.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
