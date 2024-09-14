package types

type Autoscaling_GroupInstanceRefreshPreferences struct {
	// Behavior when encountering instances in the `Standby` state in are found. Available behaviors are `Terminate`, `Ignore`, and `Wait`. Default is `Ignore`.
	StandbyInstances string `json:"standbyInstances,omitempty" yaml:"standbyInstances,omitempty"`

	// Automatically rollback if instance refresh fails. Defaults to `false`. This option may only be set to `true` when specifying a `launch_template` or `mixed_instances_policy`.
	AutoRollback bool `json:"autoRollback,omitempty" yaml:"autoRollback,omitempty"`

	// List of percentages for each checkpoint. Values must be unique and in ascending order. To replace all instances, the final number must be `100`.
	CheckpointPercentages []int `json:"checkpointPercentages,omitempty" yaml:"checkpointPercentages,omitempty"`

	// Number of seconds until a newly launched instance is configured and ready to use. Default behavior is to use the Auto Scaling Group's health check grace period.
	InstanceWarmup string `json:"instanceWarmup,omitempty" yaml:"instanceWarmup,omitempty"`

	// Amount of capacity in the Auto Scaling group that must remain healthy during an instance refresh to allow the operation to continue, as a percentage of the desired capacity of the Auto Scaling group. Defaults to `90`.
	MinHealthyPercentage int `json:"minHealthyPercentage,omitempty" yaml:"minHealthyPercentage,omitempty"`

	// Replace instances that already have your desired configuration. Defaults to `false`.
	SkipMatching bool `json:"skipMatching,omitempty" yaml:"skipMatching,omitempty"`

	// Alarm Specification for Instance Refresh.
	AlarmSpecification Autoscaling_GroupInstanceRefreshPreferencesAlarmSpecification `json:"alarmSpecification,omitempty" yaml:"alarmSpecification,omitempty"`

	// Number of seconds to wait after a checkpoint. Defaults to `3600`.
	CheckpointDelay string `json:"checkpointDelay,omitempty" yaml:"checkpointDelay,omitempty"`

	// Amount of capacity in the Auto Scaling group that can be in service and healthy, or pending, to support your workload when an instance refresh is in place, as a percentage of the desired capacity of the Auto Scaling group. Values must be between `100` and `200`, defaults to `100`.
	MaxHealthyPercentage int `json:"maxHealthyPercentage,omitempty" yaml:"maxHealthyPercentage,omitempty"`

	// Behavior when encountering instances protected from scale in are found. Available behaviors are `Refresh`, `Ignore`, and `Wait`. Default is `Ignore`.
	ScaleInProtectedInstances string `json:"scaleInProtectedInstances,omitempty" yaml:"scaleInProtectedInstances,omitempty"`
}
