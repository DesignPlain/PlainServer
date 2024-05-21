package types

type Compute_InstanceGroupManagerStatusAllInstancesConfig struct {
	// A bit indicating whether this configuration has been applied to all managed instances in the group.
	Effective bool `json:"effective,omitempty" yaml:"effective,omitempty"`
}
