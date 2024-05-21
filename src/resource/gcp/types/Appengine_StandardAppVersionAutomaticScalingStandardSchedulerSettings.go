package types

type Appengine_StandardAppVersionAutomaticScalingStandardSchedulerSettings struct {
	// Maximum number of instances to run for this version. Set to zero to disable maxInstances configuration.
	MaxInstances int `json:"maxInstances,omitempty" yaml:"maxInstances,omitempty"`

	// Minimum number of instances to run for this version. Set to zero to disable minInstances configuration.
	MinInstances int `json:"minInstances,omitempty" yaml:"minInstances,omitempty"`

	// Target CPU utilization ratio to maintain when scaling. Should be a value in the range [0.50, 0.95], zero, or a negative value.
	TargetCpuUtilization float64 `json:"targetCpuUtilization,omitempty" yaml:"targetCpuUtilization,omitempty"`

	// Target throughput utilization ratio to maintain when scaling. Should be a value in the range [0.50, 0.95], zero, or a negative value.
	TargetThroughputUtilization float64 `json:"targetThroughputUtilization,omitempty" yaml:"targetThroughputUtilization,omitempty"`
}
