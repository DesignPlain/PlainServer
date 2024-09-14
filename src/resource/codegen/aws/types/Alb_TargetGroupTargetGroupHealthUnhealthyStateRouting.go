package types

type Alb_TargetGroupTargetGroupHealthUnhealthyStateRouting struct {
	// The minimum number of targets that must be healthy. If the number of healthy targets is below this value, send traffic to all targets, including unhealthy targets. The possible values are `1` to the maximum number of targets. The default is `1`.
	MinimumHealthyTargetsCount int `json:"minimumHealthyTargetsCount,omitempty" yaml:"minimumHealthyTargetsCount,omitempty"`

	// The minimum percentage of targets that must be healthy. If the percentage of healthy targets is below this value, send traffic to all targets, including unhealthy targets. The possible values are `off` or an integer from `1` to `100`. The default is `off`.
	MinimumHealthyTargetsPercentage string `json:"minimumHealthyTargetsPercentage,omitempty" yaml:"minimumHealthyTargetsPercentage,omitempty"`
}
