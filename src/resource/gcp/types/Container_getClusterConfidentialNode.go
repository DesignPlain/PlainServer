package types

type Container_getClusterConfidentialNode struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this cluster.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
