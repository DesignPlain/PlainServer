package types

type Spanner_getInstanceAutoscalingConfigAutoscalingTarget struct {
	/*
	   Specifies the target high priority cpu utilization percentage that the autoscaler
	   should be trying to achieve for the instance.
	   This number is on a scale from 0 (no utilization) to 100 (full utilization)..
	*/
	HighPriorityCpuUtilizationPercent int `json:"highPriorityCpuUtilizationPercent,omitempty" yaml:"highPriorityCpuUtilizationPercent,omitempty"`

	/*
	   Specifies the target storage utilization percentage that the autoscaler
	   should be trying to achieve for the instance.
	   This number is on a scale from 0 (no utilization) to 100 (full utilization).
	*/
	StorageUtilizationPercent int `json:"storageUtilizationPercent,omitempty" yaml:"storageUtilizationPercent,omitempty"`
}
