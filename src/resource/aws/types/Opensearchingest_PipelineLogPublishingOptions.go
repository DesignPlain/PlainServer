package types

type Opensearchingest_PipelineLogPublishingOptions struct {
	// The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch Logs. This parameter is required if IsLoggingEnabled is set to true. See `cloudwatch_log_destination` below.
	CloudwatchLogDestination Opensearchingest_PipelineLogPublishingOptionsCloudwatchLogDestination `json:"cloudwatchLogDestination,omitempty" yaml:"cloudwatchLogDestination,omitempty"`

	// Whether logs should be published.
	IsLoggingEnabled bool `json:"isLoggingEnabled,omitempty" yaml:"isLoggingEnabled,omitempty"`
}
