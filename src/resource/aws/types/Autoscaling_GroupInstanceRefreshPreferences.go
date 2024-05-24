package types

type Autoscaling_GroupInstanceRefreshPreferences struct {
	// Automatically rollback if instance refresh fails. Defaults to `false`. This option may only be set to `true` when specifying a `launch_template` or `mixed_instances_policy`.
	AutoRollback bool `json:"autoRollback,omitempty" yaml:"autoRollback,omitempty"`

	// Number of seconds to wait after a checkpoint. Defaults to `3600`.
	CheckpointDelay string `json:"checkpointDelay,omitempty" yaml:"checkpointDelay,omitempty"`

	// Number of seconds until a newly launched instance is configured and ready to use. Default behavior is to use the Auto Scaling Group's health check grace period.
	InstanceWarmup string `json:"instanceWarmup,omitempty" yaml:"instanceWarmup,omitempty"`

	// Specifies the upper limit on the number of instances that are in the InService or Pending state with a healthy status during an instance replacement activity.
	MaxHealthyPercentage int `json:"maxHealthyPercentage,omitempty" yaml:"maxHealthyPercentage,omitempty"`

	// Behavior when encountering instances in the `Standby` state in are found. Available behaviors are `Terminate`, `Ignore`, and `Wait`. Default is `Ignore`.
	StandbyInstances string `json:"standbyInstances,omitempty" yaml:"standbyInstances,omitempty"`

	// List of percentages for each checkpoint. Values must be unique and in ascending order. To replace all instances, the final number must be `100`.
	CheckpointPercentages []int `json:"checkpointPercentages,omitempty" yaml:"checkpointPercentages,omitempty"`

	// Specifies the lower limit on the number of instances that must be in the InService state with a healthy status during an instance replacement activity.
	MinHealthyPercentage int `json:"minHealthyPercentage,omitempty" yaml:"minHealthyPercentage,omitempty"`

	// Behavior when encountering instances protected from scale in are found. Available behaviors are `Refresh`, `Ignore`, and `Wait`. Default is `Ignore`.
	ScaleInProtectedInstances string `json:"scaleInProtectedInstances,omitempty" yaml:"scaleInProtectedInstances,omitempty"`

	// Replace instances that already have your desired configuration. Defaults to `false`.
	SkipMatching bool `json:"skipMatching,omitempty" yaml:"skipMatching,omitempty"`
}
