package types

type Autoscaling_GroupMixedInstancesPolicy struct {
	// Nested argument containing settings on how to mix on-demand and Spot instances in the Auto Scaling group. Defined below.
	InstancesDistribution Autoscaling_GroupMixedInstancesPolicyInstancesDistribution `json:"instancesDistribution,omitempty" yaml:"instancesDistribution,omitempty"`

	// Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
	LaunchTemplate Autoscaling_GroupMixedInstancesPolicyLaunchTemplate `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`
}
