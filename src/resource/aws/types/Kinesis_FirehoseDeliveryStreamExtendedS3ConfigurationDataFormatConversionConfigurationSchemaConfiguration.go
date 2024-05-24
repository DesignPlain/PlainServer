package types

type Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationSchemaConfiguration struct {
	// If you don't specify an AWS Region, the default is the current region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The role that Kinesis Data Firehose can use to access AWS Glue. This role must be in the same account you use for Kinesis Data Firehose. Cross-account roles aren't allowed.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Specifies the AWS Glue table that contains the column information that constitutes your data schema.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// Specifies the table version for the output data schema. Defaults to `LATEST`.
	VersionId string `json:"versionId,omitempty" yaml:"versionId,omitempty"`

	// The ID of the AWS Glue Data Catalog. If you don't supply this, the AWS account ID is used by default.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Specifies the name of the AWS Glue database that contains the schema for the output data.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`
}
