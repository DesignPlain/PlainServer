package types

type Container_ClusterIdentityServiceConfig struct {
	// Whether to enable the Identity Service component. It is disabled by default. Set `enabled=true` to enable.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
