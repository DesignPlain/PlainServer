package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfig struct {
	// Cloud Run runtime configuration.
	CloudRun Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigCloudRun `json:"cloudRun,omitempty" yaml:"cloudRun,omitempty"`

	// Kubernetes runtime configuration.
	Kubernetes Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigKubernetes `json:"kubernetes,omitempty" yaml:"kubernetes,omitempty"`
}
