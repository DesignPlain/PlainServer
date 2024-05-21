package types

type Appengine_FlexibleAppVersionAutomaticScalingCpuUtilization struct {
	// Period of time over which CPU utilization is calculated.
	AggregationWindowLength string `json:"aggregationWindowLength,omitempty" yaml:"aggregationWindowLength,omitempty"`

	// Target CPU utilization ratio to maintain when scaling. Must be between 0 and 1.
	TargetUtilization float64 `json:"targetUtilization,omitempty" yaml:"targetUtilization,omitempty"`
}
