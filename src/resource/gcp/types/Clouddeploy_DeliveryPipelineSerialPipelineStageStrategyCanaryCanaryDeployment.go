package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCanaryDeployment struct {
	// Required. The percentage based deployments that will occur as a part of a `Rollout`. List is expected in ascending order and each integer n is 0 <= n < 100.
	Percentages []int `json:"percentages,omitempty" yaml:"percentages,omitempty"`

	// Optional. Configuration for the postdeploy job of the last phase. If this is not configured, postdeploy job will not be present.
	Postdeploy Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCanaryDeploymentPostdeploy `json:"postdeploy,omitempty" yaml:"postdeploy,omitempty"`

	// Optional. Configuration for the predeploy job of the first phase. If this is not configured, predeploy job will not be present.
	Predeploy Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCanaryDeploymentPredeploy `json:"predeploy,omitempty" yaml:"predeploy,omitempty"`

	// Whether to run verify tests after each percentage deployment.
	Verify bool `json:"verify,omitempty" yaml:"verify,omitempty"`
}
