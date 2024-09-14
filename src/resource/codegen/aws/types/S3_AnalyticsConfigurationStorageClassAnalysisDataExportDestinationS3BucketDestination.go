package types

type S3_AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestination struct {
	// Account ID that owns the destination bucket.
	BucketAccountId string `json:"bucketAccountId,omitempty" yaml:"bucketAccountId,omitempty"`

	// ARN of the destination bucket.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// Output format of exported analytics data. Allowed values: `CSV`. Default value: `CSV`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// Prefix to append to exported analytics data.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
