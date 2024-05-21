package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanary struct {
	// Configures the progressive based deployment for a Target.
	CanaryDeployment Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCanaryDeployment `json:"canaryDeployment,omitempty" yaml:"canaryDeployment,omitempty"`

	// Configures the progressive based deployment for a Target, but allows customizing at the phase level where a phase represents each of the percentage deployments.
	CustomCanaryDeployment Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCustomCanaryDeployment `json:"customCanaryDeployment,omitempty" yaml:"customCanaryDeployment,omitempty"`

	// Optional. Runtime specific configurations for the deployment strategy. The runtime configuration is used to determine how Cloud Deploy will split traffic to enable a progressive deployment.
	RuntimeConfig Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfig `json:"runtimeConfig,omitempty" yaml:"runtimeConfig,omitempty"`
}
