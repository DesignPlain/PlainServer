package types

type Kinesis_AnalyticsApplicationReferenceDataSources struct {
	// The ARN of the Kinesis Analytics Application.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The S3 configuration for the reference data source. See S3 Reference below for more details.
	S3 Kinesis_AnalyticsApplicationReferenceDataSourcesS3 `json:"s3,omitempty" yaml:"s3,omitempty"`

	// The Schema format of the data in the streaming source. See Source Schema below for more details.
	Schema Kinesis_AnalyticsApplicationReferenceDataSourcesSchema `json:"schema,omitempty" yaml:"schema,omitempty"`

	// The in-application Table Name.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
