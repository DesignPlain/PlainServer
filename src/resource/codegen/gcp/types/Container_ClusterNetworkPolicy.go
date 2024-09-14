package types

type Container_ClusterNetworkPolicy struct {
	// Whether network policy is enabled on the cluster.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED.
	Provider string `json:"provider,omitempty" yaml:"provider,omitempty"`
}
