package types

type Opsworks_CustomLayerLoadBasedAutoScalingUpscaling struct {
	// The amount of time (in minutes) after a scaling event occurs that AWS OpsWorks Stacks should ignore metrics and suppress additional scaling events.
	IgnoreMetricsTime int `json:"ignoreMetricsTime,omitempty" yaml:"ignoreMetricsTime,omitempty"`

	// The number of instances to add or remove when the load exceeds a threshold.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	// The load threshold. A value of -1 disables the threshold.
	LoadThreshold float64 `json:"loadThreshold,omitempty" yaml:"loadThreshold,omitempty"`

	// The memory utilization threshold, as a percent of the available memory. A value of -1 disables the threshold.
	MemoryThreshold float64 `json:"memoryThreshold,omitempty" yaml:"memoryThreshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold before more instances are added or removed.
	ThresholdsWaitTime int `json:"thresholdsWaitTime,omitempty" yaml:"thresholdsWaitTime,omitempty"`

	// Custom Cloudwatch auto scaling alarms, to be used as thresholds. This parameter takes a list of up to five alarm names, which are case sensitive and must be in the same region as the stack.
	Alarms []string `json:"alarms,omitempty" yaml:"alarms,omitempty"`

	// The CPU utilization threshold, as a percent of the available CPU. A value of -1 disables the threshold.
	CpuThreshold float64 `json:"cpuThreshold,omitempty" yaml:"cpuThreshold,omitempty"`
}
