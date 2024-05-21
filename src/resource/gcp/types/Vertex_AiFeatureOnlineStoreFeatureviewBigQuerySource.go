package types

type Vertex_AiFeatureOnlineStoreFeatureviewBigQuerySource struct {
	// Columns to construct entityId / row keys. Start by supporting 1 only.
	EntityIdColumns []string `json:"entityIdColumns,omitempty" yaml:"entityIdColumns,omitempty"`

	// The BigQuery view URI that will be materialized on each sync trigger based on FeatureView.SyncConfig.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
