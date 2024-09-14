package types

type Integrationconnectors_ConnectionNodeConfig struct {
	// Minimum number of nodes in the runtime nodes.
	MaxNodeCount int `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`

	// Minimum number of nodes in the runtime nodes.
	MinNodeCount int `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`
}
