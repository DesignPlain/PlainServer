package types

type Glue_MLTransformInputRecordTable struct {
	// A unique identifier for the AWS Glue Data Catalog.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// The name of the connection to the AWS Glue Data Catalog.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// A database name in the AWS Glue Data Catalog.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// A table name in the AWS Glue Data Catalog.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
