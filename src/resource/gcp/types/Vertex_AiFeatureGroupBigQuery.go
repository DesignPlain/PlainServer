package types

type Vertex_AiFeatureGroupBigQuery struct {
	/*
	   The BigQuery source URI that points to either a BigQuery Table or View.
	   Structure is documented below.
	*/
	BigQuerySource Vertex_AiFeatureGroupBigQueryBigQuerySource `json:"bigQuerySource,omitempty" yaml:"bigQuerySource,omitempty"`

	// Columns to construct entityId / row keys. Currently only supports 1 entity_id_column. If not provided defaults to entityId.
	EntityIdColumns []string `json:"entityIdColumns,omitempty" yaml:"entityIdColumns,omitempty"`
}
