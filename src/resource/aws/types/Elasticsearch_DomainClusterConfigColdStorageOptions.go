package types

type Elasticsearch_DomainClusterConfigColdStorageOptions struct {
	// Boolean to enable cold storage for an Elasticsearch domain. Defaults to `false`. Master and ultrawarm nodes must be enabled for cold storage.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
