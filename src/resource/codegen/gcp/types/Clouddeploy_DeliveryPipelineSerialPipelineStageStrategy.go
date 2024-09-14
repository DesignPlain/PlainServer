package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategy struct {
	// Canary deployment strategy provides progressive percentage based deployments to a Target.
	Canary Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanary `json:"canary,omitempty" yaml:"canary,omitempty"`

	// Standard deployment strategy executes a single deploy and allows verifying the deployment.
	Standard Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyStandard `json:"standard,omitempty" yaml:"standard,omitempty"`
}
