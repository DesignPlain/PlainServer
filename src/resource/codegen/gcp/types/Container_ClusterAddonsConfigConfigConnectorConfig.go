package types

type Container_ClusterAddonsConfigConfigConnectorConfig struct {
	// Enable Binary Authorization for this cluster. Deprecated in favor of `evaluation_mode`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
