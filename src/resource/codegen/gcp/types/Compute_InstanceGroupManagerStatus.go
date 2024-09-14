package types

type Compute_InstanceGroupManagerStatus struct {
	/*
	   Properties to set on all instances in the group. After setting
	   allInstancesConfig on the group, you must update the group's instances to
	   apply the configuration.
	*/
	AllInstancesConfigs []Compute_InstanceGroupManagerStatusAllInstancesConfig `json:"allInstancesConfigs,omitempty" yaml:"allInstancesConfigs,omitempty"`

	// A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
	IsStable bool `json:"isStable,omitempty" yaml:"isStable,omitempty"`

	// Stateful status of the given Instance Group Manager.
	Statefuls []Compute_InstanceGroupManagerStatusStateful `json:"statefuls,omitempty" yaml:"statefuls,omitempty"`

	// A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances' target version are specified by version field on Instance Group Manager.
	VersionTargets []Compute_InstanceGroupManagerStatusVersionTarget `json:"versionTargets,omitempty" yaml:"versionTargets,omitempty"`
}
