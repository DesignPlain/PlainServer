package types

type Container_ClusterClusterAutoscaling struct {
	/*
	   Whether node auto-provisioning is enabled. Must be supplied for GKE Standard clusters, `true` is implied
	   for autopilot clusters. Resource limits for `cpu` and `memory` must be defined to enable node auto-provisioning for GKE Standard.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   Global constraints for machine resources in the
	   cluster. Configuring the `cpu` and `memory` types is required if node
	   auto-provisioning is enabled. These limits will apply to node pool autoscaling
	   in addition to node auto-provisioning. Structure is documented below.
	*/
	ResourceLimits []Container_ClusterClusterAutoscalingResourceLimit `json:"resourceLimits,omitempty" yaml:"resourceLimits,omitempty"`

	/*
	   Contains defaults for a node pool created by NAP. A subset of fields also apply to
	   GKE Autopilot clusters.
	   Structure is documented below.
	*/
	AutoProvisioningDefaults Container_ClusterClusterAutoscalingAutoProvisioningDefaults `json:"autoProvisioningDefaults,omitempty" yaml:"autoProvisioningDefaults,omitempty"`

	/*
	   Configuration
	   options for the [Autoscaling profile](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler#autoscaling_profiles)
	   feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability
	   when deciding to remove nodes from a cluster. Can be `BALANCED` or `OPTIMIZE_UTILIZATION`. Defaults to `BALANCED`.
	*/
	AutoscalingProfile string `json:"autoscalingProfile,omitempty" yaml:"autoscalingProfile,omitempty"`
}
