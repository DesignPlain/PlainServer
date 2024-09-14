package types

type S3_AnalyticsConfigurationStorageClassAnalysisDataExportDestination struct {
	// Analytics data export currently only supports an S3 bucket destination (documented below).
	S3BucketDestination S3_AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestination `json:"s3BucketDestination,omitempty" yaml:"s3BucketDestination,omitempty"`
}
