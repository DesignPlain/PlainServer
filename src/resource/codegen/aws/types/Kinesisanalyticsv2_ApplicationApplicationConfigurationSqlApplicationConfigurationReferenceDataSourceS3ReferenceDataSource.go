package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSource struct {
	// The ARN of the S3 bucket.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// The object key name containing the reference data.
	FileKey string `json:"fileKey,omitempty" yaml:"fileKey,omitempty"`
}
