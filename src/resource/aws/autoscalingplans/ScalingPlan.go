package autoscalingplans

import types "DesignSphere_Server/src/resource/aws/types"

type ScalingPlan struct {
	// CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource types.Autoscalingplans_ScalingPlanApplicationSource `json:"applicationSource,omitempty" yaml:"applicationSource,omitempty"`

	// Name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
	ScalingInstructions []types.Autoscalingplans_ScalingPlanScalingInstruction `json:"scalingInstructions,omitempty" yaml:"scalingInstructions,omitempty"`
}
