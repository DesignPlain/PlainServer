package types

type Kinesis_AnalyticsApplicationReferenceDataSourcesS3 struct {
	// The S3 Bucket ARN.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// The File Key name containing reference data.
	FileKey string `json:"fileKey,omitempty" yaml:"fileKey,omitempty"`

	// The IAM Role ARN to read the data.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
