package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCustomCanaryDeploymentPhaseConfig struct {
	// Optional. Configuration for the predeploy job of this phase. If this is not configured, predeploy job will not be present for this phase.
	Predeploy Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCustomCanaryDeploymentPhaseConfigPredeploy `json:"predeploy,omitempty" yaml:"predeploy,omitempty"`

	// Skaffold profiles to use when rendering the manifest for this phase. These are in addition to the profiles list specified in the `DeliveryPipeline` stage.
	Profiles []string `json:"profiles,omitempty" yaml:"profiles,omitempty"`

	/*
	   Whether to run verify tests after the deployment.

	   - - -
	*/
	Verify bool `json:"verify,omitempty" yaml:"verify,omitempty"`

	// Required. Percentage deployment for the phase.
	Percentage int `json:"percentage,omitempty" yaml:"percentage,omitempty"`

	// Required. The ID to assign to the `Rollout` phase. This value must consist of lower-case letters, numbers, and hyphens, start with a letter and end with a letter or a number, and have a max length of 63 characters. In other words, it must match the following regex: `^a-z?$`.
	PhaseId string `json:"phaseId,omitempty" yaml:"phaseId,omitempty"`

	// Optional. Configuration for the postdeploy job of this phase. If this is not configured, postdeploy job will not be present for this phase.
	Postdeploy Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCustomCanaryDeploymentPhaseConfigPostdeploy `json:"postdeploy,omitempty" yaml:"postdeploy,omitempty"`
}
