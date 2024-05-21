package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigKubernetesServiceNetworking struct {
	// Required. Name of the Kubernetes Deployment whose traffic is managed by the specified Service.
	Deployment string `json:"deployment,omitempty" yaml:"deployment,omitempty"`

	// Optional. Whether to disable Pod overprovisioning. If Pod overprovisioning is disabled then Cloud Deploy will limit the number of total Pods used for the deployment strategy to the number of Pods the Deployment has on the cluster.
	DisablePodOverprovisioning bool `json:"disablePodOverprovisioning,omitempty" yaml:"disablePodOverprovisioning,omitempty"`

	// Required. Name of the Kubernetes Service.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
