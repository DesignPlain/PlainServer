package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCanaryDeploymentPostdeploy struct {
	// Optional. A sequence of skaffold custom actions to invoke during execution of the postdeploy job.
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`
}
