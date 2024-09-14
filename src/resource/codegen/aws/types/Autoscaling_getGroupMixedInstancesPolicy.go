package types

type Autoscaling_getGroupMixedInstancesPolicy struct {
	// List of instances distribution objects.
	InstancesDistributions []Autoscaling_getGroupMixedInstancesPolicyInstancesDistribution `json:"instancesDistributions,omitempty" yaml:"instancesDistributions,omitempty"`

	// List of launch templates along with the overrides.
	LaunchTemplates []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplate `json:"launchTemplates,omitempty" yaml:"launchTemplates,omitempty"`
}
