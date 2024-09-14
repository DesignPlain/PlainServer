package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig struct {
	// The desired file size, in MB, for each output file that Amazon AppFlow writes to the flow destination. Integer value.
	TargetFileSize int `json:"targetFileSize,omitempty" yaml:"targetFileSize,omitempty"`

	// Whether Amazon AppFlow aggregates the flow records into a single file, or leave them unaggregated. Valid values are `None` and `SingleFile`.
	AggregationType string `json:"aggregationType,omitempty" yaml:"aggregationType,omitempty"`
}
