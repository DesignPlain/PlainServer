package appfabric

import types "libds/aws/types"

type IngestionDestination struct {
	// Contains information about the destination of ingested data.
	DestinationConfiguration types.Appfabric_IngestionDestinationDestinationConfiguration `json:"destinationConfiguration,omitempty" yaml:"destinationConfiguration,omitempty"`

	// The Amazon Resource Name (ARN) of the ingestion to use for the request.
	IngestionArn string `json:"ingestionArn,omitempty" yaml:"ingestionArn,omitempty"`

	// Contains information about how ingested data is processed.
	ProcessingConfiguration types.Appfabric_IngestionDestinationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Appfabric_IngestionDestinationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The Amazon Resource Name (ARN) of the app bundle to use for the request.
	AppBundleArn string `json:"appBundleArn,omitempty" yaml:"appBundleArn,omitempty"`
}
