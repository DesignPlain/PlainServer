package types

type Mskconnect_ConnectorCapacityAutoscalingScaleOutPolicy struct {
	// The CPU utilization percentage threshold at which you want connector scale out to be triggered.
	CpuUtilizationPercentage int `json:"cpuUtilizationPercentage,omitempty" yaml:"cpuUtilizationPercentage,omitempty"`
}
