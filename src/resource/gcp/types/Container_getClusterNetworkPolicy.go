package types

type Container_getClusterNetworkPolicy struct {
	// Whether network policy is enabled on the cluster.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The selected network policy provider.
	Provider string `json:"provider,omitempty" yaml:"provider,omitempty"`
}
