package types

type Dax_ClusterServerSideEncryption struct {
	// Whether to enable encryption at rest. Defaults to `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
