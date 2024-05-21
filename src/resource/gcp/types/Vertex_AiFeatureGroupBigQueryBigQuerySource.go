package types

type Vertex_AiFeatureGroupBigQueryBigQuerySource struct {
	// BigQuery URI to a table, up to 2000 characters long. For example: `bq://projectId.bqDatasetId.bqTableId.`
	InputUri string `json:"inputUri,omitempty" yaml:"inputUri,omitempty"`
}
