package types

type Compute_InstanceGroupManagerStatusStatefulPerInstanceConfig struct {
	// A bit indicating if all of the group's per-instance configs (listed in the output of a listPerInstanceConfigs API call) have status `EFFECTIVE` or there are no per-instance-configs.
	AllEffective bool `json:"allEffective,omitempty" yaml:"allEffective,omitempty"`
}
