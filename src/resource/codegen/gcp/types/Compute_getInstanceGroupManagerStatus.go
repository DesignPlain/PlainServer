package types

type Compute_getInstanceGroupManagerStatus struct {
	// Status of all-instances configuration on the group.
	AllInstancesConfigs []Compute_getInstanceGroupManagerStatusAllInstancesConfig `json:"allInstancesConfigs,omitempty" yaml:"allInstancesConfigs,omitempty"`

	// A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
	IsStable bool `json:"isStable,omitempty" yaml:"isStable,omitempty"`

	// Stateful status of the given Instance Group Manager.
	Statefuls []Compute_getInstanceGroupManagerStatusStateful `json:"statefuls,omitempty" yaml:"statefuls,omitempty"`

	// A status of consistency of Instances' versions with their target version specified by version field on Instance Group Manager.
	VersionTargets []Compute_getInstanceGroupManagerStatusVersionTarget `json:"versionTargets,omitempty" yaml:"versionTargets,omitempty"`
}
