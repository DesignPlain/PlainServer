package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSource struct {
	// The name of the in-application table to create.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	//
	ReferenceId string `json:"referenceId,omitempty" yaml:"referenceId,omitempty"`

	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
	ReferenceSchema Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchema `json:"referenceSchema,omitempty" yaml:"referenceSchema,omitempty"`

	// Identifies the S3 bucket and object that contains the reference data.
	S3ReferenceDataSource Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSource `json:"s3ReferenceDataSource,omitempty" yaml:"s3ReferenceDataSource,omitempty"`
}
