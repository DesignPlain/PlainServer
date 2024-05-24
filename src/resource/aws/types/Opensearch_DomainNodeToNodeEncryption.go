package types

type Opensearch_DomainNodeToNodeEncryption struct {
	// Whether to enable node-to-node encryption. If the `node_to_node_encryption` block is not provided then this defaults to `false`. Enabling node-to-node encryption of a new domain requires an `engine_version` of `OpenSearch_X.Y` or `Elasticsearch_6.0` or greater.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
