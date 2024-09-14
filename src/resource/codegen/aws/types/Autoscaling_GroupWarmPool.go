package types

type Autoscaling_GroupWarmPool struct {
	// Total maximum number of instances that are allowed to be in the warm pool or in any state except Terminated for the Auto Scaling group.
	MaxGroupPreparedCapacity int `json:"maxGroupPreparedCapacity,omitempty" yaml:"maxGroupPreparedCapacity,omitempty"`

	// Minimum number of instances to maintain in the warm pool. This helps you to ensure that there is always a certain number of warmed instances available to handle traffic spikes. Defaults to 0 if not specified.
	MinSize int `json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// Sets the instance state to transition to after the lifecycle hooks finish. Valid values are: Stopped (default), Running or Hibernated.
	PoolState string `json:"poolState,omitempty" yaml:"poolState,omitempty"`

	// Whether instances in the Auto Scaling group can be returned to the warm pool on scale in. The default is to terminate instances in the Auto Scaling group when the group scales in.
	InstanceReusePolicy Autoscaling_GroupWarmPoolInstanceReusePolicy `json:"instanceReusePolicy,omitempty" yaml:"instanceReusePolicy,omitempty"`
}
