package types

type Opsworks_JavaAppLayerLoadBasedAutoScalingDownscaling struct {
	//
	CpuThreshold float64 `json:"cpuThreshold,omitempty" yaml:"cpuThreshold,omitempty"`

	//
	IgnoreMetricsTime int `json:"ignoreMetricsTime,omitempty" yaml:"ignoreMetricsTime,omitempty"`

	//
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	//
	LoadThreshold float64 `json:"loadThreshold,omitempty" yaml:"loadThreshold,omitempty"`

	//
	MemoryThreshold float64 `json:"memoryThreshold,omitempty" yaml:"memoryThreshold,omitempty"`

	//
	ThresholdsWaitTime int `json:"thresholdsWaitTime,omitempty" yaml:"thresholdsWaitTime,omitempty"`

	//
	Alarms []string `json:"alarms,omitempty" yaml:"alarms,omitempty"`
}
