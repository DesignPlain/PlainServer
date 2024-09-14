package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfig struct {
	// Whether Amazon AppFlow aggregates the flow records into a single file, or leave them unaggregated. Valid values are `None` and `SingleFile`.
	AggregationType string `json:"aggregationType,omitempty" yaml:"aggregationType,omitempty"`
}
