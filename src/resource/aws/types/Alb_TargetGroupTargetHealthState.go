package types

type Alb_TargetGroupTargetHealthState struct {
	// Indicates whether the load balancer terminates connections to unhealthy targets. Possible values are `true` or `false`. Default: `true`.
	EnableUnhealthyConnectionTermination bool `json:"enableUnhealthyConnectionTermination,omitempty" yaml:"enableUnhealthyConnectionTermination,omitempty"`
}
