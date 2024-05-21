package types

type Clouddeploy_DeliveryPipelineCondition struct {
	// Details on the whether the targets enumerated in the pipeline are of the same type.
	TargetsTypeConditions []Clouddeploy_DeliveryPipelineConditionTargetsTypeCondition `json:"targetsTypeConditions,omitempty" yaml:"targetsTypeConditions,omitempty"`

	// Details around the Pipeline's overall status.
	PipelineReadyConditions []Clouddeploy_DeliveryPipelineConditionPipelineReadyCondition `json:"pipelineReadyConditions,omitempty" yaml:"pipelineReadyConditions,omitempty"`

	// Details around targets enumerated in the pipeline.
	TargetsPresentConditions []Clouddeploy_DeliveryPipelineConditionTargetsPresentCondition `json:"targetsPresentConditions,omitempty" yaml:"targetsPresentConditions,omitempty"`
}
