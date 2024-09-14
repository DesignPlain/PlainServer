package types

type Appflow_FlowMetadataCatalogConfigGlueDataCatalog struct {
	// The ARN of an IAM role that grants AppFlow the permissions it needs to create Data Catalog tables, databases, and partitions.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A naming prefix for each Data Catalog table that Amazon AppFlow creates
	TablePrefix string `json:"tablePrefix,omitempty" yaml:"tablePrefix,omitempty"`

	// The name of an existing Glue database to store the metadata tables that Amazon AppFlow creates.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`
}
