package types

type Glue_DataQualityRulesetTargetTable struct {
	// The catalog id where the AWS Glue table exists.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Name of the database where the AWS Glue table exists.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Name of the AWS Glue table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
