package types

type Kinesis_AnalyticsApplicationReferenceDataSourcesS3 struct {
	// The S3 Bucket ARN.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// The File Key name containing reference data.
	FileKey string `json:"fileKey,omitempty" yaml:"fileKey,omitempty"`

	// The ARN of the IAM Role used to send application messages.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
