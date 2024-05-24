package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSource struct {
	// The ARN for the S3 bucket containing the application code.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// The file key for the object containing the application code.
	FileKey string `json:"fileKey,omitempty" yaml:"fileKey,omitempty"`
}
