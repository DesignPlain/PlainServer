package types

type Container_getClusterClusterAutoscaling struct {
	// Global constraints for machine resources in the cluster. Configuring the cpu and memory types is required if node auto-provisioning is enabled. These limits will apply to node pool autoscaling in addition to node auto-provisioning.
	ResourceLimits []Container_getClusterClusterAutoscalingResourceLimit `json:"resourceLimits,omitempty" yaml:"resourceLimits,omitempty"`

	// Contains defaults for a node pool created by NAP.
	AutoProvisioningDefaults []Container_getClusterClusterAutoscalingAutoProvisioningDefault `json:"autoProvisioningDefaults,omitempty" yaml:"autoProvisioningDefaults,omitempty"`

	// Configuration options for the Autoscaling profile feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability when deciding to remove nodes from a cluster. Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults to BALANCED.
	AutoscalingProfile string `json:"autoscalingProfile,omitempty" yaml:"autoscalingProfile,omitempty"`

	// Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
