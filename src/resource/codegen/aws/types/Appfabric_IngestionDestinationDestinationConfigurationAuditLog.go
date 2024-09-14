package types

type Appfabric_IngestionDestinationDestinationConfigurationAuditLog struct {
	// Contains information about an audit log destination. Only one destination (Firehose Stream) or (S3 Bucket) can be specified.
	Destination Appfabric_IngestionDestinationDestinationConfigurationAuditLogDestination `json:"destination,omitempty" yaml:"destination,omitempty"`
}
