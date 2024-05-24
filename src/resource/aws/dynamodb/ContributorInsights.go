package dynamodb

type ContributorInsights struct {
	// The name of the table to enable contributor insights
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// The global secondary index name
	IndexName string `json:"indexName,omitempty" yaml:"indexName,omitempty"`
}
