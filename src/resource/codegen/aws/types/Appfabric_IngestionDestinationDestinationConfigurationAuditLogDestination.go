package types

type Appfabric_IngestionDestinationDestinationConfigurationAuditLogDestination struct {
	// Contains information about an Amazon Data Firehose delivery stream.
	FirehoseStream Appfabric_IngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStream `json:"firehoseStream,omitempty" yaml:"firehoseStream,omitempty"`

	// Contains information about an Amazon S3 bucket.
	S3Bucket Appfabric_IngestionDestinationDestinationConfigurationAuditLogDestinationS3Bucket `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`
}
