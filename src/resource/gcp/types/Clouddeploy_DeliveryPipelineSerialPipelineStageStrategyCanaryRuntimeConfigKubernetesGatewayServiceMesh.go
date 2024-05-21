package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh struct {
	// Required. Name of the Kubernetes Service.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// Optional. The amount of time to migrate traffic back from the canary Service to the original Service during the stable phase deployment. If specified, must be between 15s and 3600s. If unspecified, there is no cutback time.
	StableCutbackDuration string `json:"stableCutbackDuration,omitempty" yaml:"stableCutbackDuration,omitempty"`

	// Required. Name of the Kubernetes Deployment whose traffic is managed by the specified HTTPRoute and Service.
	Deployment string `json:"deployment,omitempty" yaml:"deployment,omitempty"`

	// Required. Name of the Gateway API HTTPRoute.
	HttpRoute string `json:"httpRoute,omitempty" yaml:"httpRoute,omitempty"`

	// Optional. The time to wait for route updates to propagate. The maximum configurable time is 3 hours, in seconds format. If unspecified, there is no wait time.
	RouteUpdateWaitTime string `json:"routeUpdateWaitTime,omitempty" yaml:"routeUpdateWaitTime,omitempty"`
}
