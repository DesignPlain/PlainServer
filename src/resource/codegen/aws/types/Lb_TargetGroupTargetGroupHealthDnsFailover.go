package types

type Lb_TargetGroupTargetGroupHealthDnsFailover struct {
	// The minimum number of targets that must be healthy. If the number of healthy targets is below this value, mark the zone as unhealthy in DNS, so that traffic is routed only to healthy zones. The possible values are `off` or an integer from `1` to the maximum number of targets. The default is `off`.
	MinimumHealthyTargetsCount string `json:"minimumHealthyTargetsCount,omitempty" yaml:"minimumHealthyTargetsCount,omitempty"`

	// The minimum percentage of targets that must be healthy. If the percentage of healthy targets is below this value, mark the zone as unhealthy in DNS, so that traffic is routed only to healthy zones. The possible values are `off` or an integer from `1` to `100`. The default is `off`.
	MinimumHealthyTargetsPercentage string `json:"minimumHealthyTargetsPercentage,omitempty" yaml:"minimumHealthyTargetsPercentage,omitempty"`
}
