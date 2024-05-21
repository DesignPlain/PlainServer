package types

type Clouddeploy_DeliveryPipelineConditionPipelineReadyCondition struct {
	// True if the Pipeline is in a valid state. Otherwise at least one condition in `PipelineCondition` is in an invalid state. Iterate over those conditions and see which condition(s) has status = false to find out what is wrong with the Pipeline.
	Status bool `json:"status,omitempty" yaml:"status,omitempty"`

	// Output only. Most recent time at which the pipeline was updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
