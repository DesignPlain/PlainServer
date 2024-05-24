package types

type Finspace_KxClusterAutoScalingConfiguration struct {
	// Duration in seconds that FinSpace will wait after a scale in event before initiating another scaling event.
	ScaleInCooldownSeconds float64 `json:"scaleInCooldownSeconds,omitempty" yaml:"scaleInCooldownSeconds,omitempty"`

	// Duration in seconds that FinSpace will wait after a scale out event before initiating another scaling event.
	ScaleOutCooldownSeconds float64 `json:"scaleOutCooldownSeconds,omitempty" yaml:"scaleOutCooldownSeconds,omitempty"`

	// Metric your cluster will track in order to scale in and out. For example, CPU_UTILIZATION_PERCENTAGE is the average CPU usage across all nodes in a cluster.
	AutoScalingMetric string `json:"autoScalingMetric,omitempty" yaml:"autoScalingMetric,omitempty"`

	// Highest number of nodes to scale. Cannot be greater than 5
	MaxNodeCount int `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`

	// Desired value of chosen `auto_scaling_metric`. When metric drops below this value, cluster will scale in. When metric goes above this value, cluster will scale out. Can be set between 0 and 100 percent.
	MetricTarget float64 `json:"metricTarget,omitempty" yaml:"metricTarget,omitempty"`

	// Lowest number of nodes to scale. Must be at least 1 and less than the `max_node_count`. If nodes in cluster belong to multiple availability zones, then `min_node_count` must be at least 3.
	MinNodeCount int `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`
}
