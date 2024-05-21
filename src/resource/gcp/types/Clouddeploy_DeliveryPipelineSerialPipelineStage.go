package types

type Clouddeploy_DeliveryPipelineSerialPipelineStage struct {
	// The target_id to which this stage points. This field refers exclusively to the last segment of a target name. For example, this field would just be `my-target` (rather than `projects/project/locations/location/targets/my-target`). The location of the `Target` is inferred to be the same as the location of the `DeliveryPipeline` that contains this `Stage`.
	TargetId string `json:"targetId,omitempty" yaml:"targetId,omitempty"`

	// Optional. The deploy parameters to use for the target in this stage.
	DeployParameters []Clouddeploy_DeliveryPipelineSerialPipelineStageDeployParameter `json:"deployParameters,omitempty" yaml:"deployParameters,omitempty"`

	// Skaffold profiles to use when rendering the manifest for this stage's `Target`.
	Profiles []string `json:"profiles,omitempty" yaml:"profiles,omitempty"`

	// Optional. The strategy to use for a `Rollout` to this stage.
	Strategy Clouddeploy_DeliveryPipelineSerialPipelineStageStrategy `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
