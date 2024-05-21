package types



type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigKubernetes struct {
	// Kubernetes Service networking configuration.
	ServiceNetworking Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigKubernetesServiceNetworking `json:"serviceNetworking,omitempty" yaml:"serviceNetworking,omitempty"`

	// Kubernetes Gateway API service mesh configuration.
	GatewayServiceMesh Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh `json:"gatewayServiceMesh,omitempty" yaml:"gatewayServiceMesh,omitempty"`
}
