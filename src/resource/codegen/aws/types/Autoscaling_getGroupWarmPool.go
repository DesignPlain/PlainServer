package types

type Autoscaling_getGroupWarmPool struct {
	// List of instance reuse policy objects.
	InstanceReusePolicies []Autoscaling_getGroupWarmPoolInstanceReusePolicy `json:"instanceReusePolicies,omitempty" yaml:"instanceReusePolicies,omitempty"`

	//
	MaxGroupPreparedCapacity int `json:"maxGroupPreparedCapacity,omitempty" yaml:"maxGroupPreparedCapacity,omitempty"`

	// Minimum number of instances to maintain in the warm pool.
	MinSize int `json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// Instance state to transition to after the lifecycle actions are complete.
	PoolState string `json:"poolState,omitempty" yaml:"poolState,omitempty"`
}
