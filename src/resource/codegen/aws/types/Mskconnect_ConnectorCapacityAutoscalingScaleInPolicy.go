package types

type Mskconnect_ConnectorCapacityAutoscalingScaleInPolicy struct {
	// Specifies the CPU utilization percentage threshold at which you want connector scale in to be triggered.
	CpuUtilizationPercentage int `json:"cpuUtilizationPercentage,omitempty" yaml:"cpuUtilizationPercentage,omitempty"`
}
