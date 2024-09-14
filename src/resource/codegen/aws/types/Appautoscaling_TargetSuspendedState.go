package types

type Appautoscaling_TargetSuspendedState struct {
	// Whether scale out by a target tracking scaling policy or a step scaling policy is suspended. Default is `false`.
	DynamicScalingOutSuspended bool `json:"dynamicScalingOutSuspended,omitempty" yaml:"dynamicScalingOutSuspended,omitempty"`

	// Whether scheduled scaling is suspended. Default is `false`.
	ScheduledScalingSuspended bool `json:"scheduledScalingSuspended,omitempty" yaml:"scheduledScalingSuspended,omitempty"`

	// Whether scale in by a target tracking scaling policy or a step scaling policy is suspended. Default is `false`.
	DynamicScalingInSuspended bool `json:"dynamicScalingInSuspended,omitempty" yaml:"dynamicScalingInSuspended,omitempty"`
}
