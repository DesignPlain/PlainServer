package types

type S3control_StorageLensConfigurationStorageLensConfigurationDataExport struct {
	// Amazon CloudWatch publishing for S3 Storage Lens metrics. See Cloud Watch Metrics below for more details.
	CloudWatchMetrics S3control_StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetrics `json:"cloudWatchMetrics,omitempty" yaml:"cloudWatchMetrics,omitempty"`

	// The bucket where the S3 Storage Lens metrics export will be located. See S3 Bucket Destination below for more details.
	S3BucketDestination S3control_StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestination `json:"s3BucketDestination,omitempty" yaml:"s3BucketDestination,omitempty"`
}
