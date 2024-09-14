package types

type Opensearch_OutboundConnectionConnectionPropertiesCrossClusterSearch struct {
	// Skips unavailable clusters and can only be used for cross-cluster searches. Accepted values are `ENABLED` or `DISABLED`.
	SkipUnavailable string `json:"skipUnavailable,omitempty" yaml:"skipUnavailable,omitempty"`
}
