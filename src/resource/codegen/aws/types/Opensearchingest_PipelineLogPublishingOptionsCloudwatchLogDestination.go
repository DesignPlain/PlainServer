package types

type Opensearchingest_PipelineLogPublishingOptionsCloudwatchLogDestination struct {
	// The name of the CloudWatch Logs group to send pipeline logs to. You can specify an existing log group or create a new one. For example, /aws/OpenSearchService/IngestionService/my-pipeline.
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`
}
