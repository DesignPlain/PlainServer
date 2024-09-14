package types

type Lb_TargetGroupTargetHealthState struct {
	// Indicates whether the load balancer terminates connections to unhealthy targets. Possible values are `true` or `false`. Default: `true`.
	EnableUnhealthyConnectionTermination bool `json:"enableUnhealthyConnectionTermination,omitempty" yaml:"enableUnhealthyConnectionTermination,omitempty"`

	// Indicates the time to wait for in-flight requests to complete when a target becomes unhealthy. The range is `0-360000`. This value has to be set only if `enable_unhealthy_connection_termination` is set to false. Default: `0`.
	UnhealthyDrainingInterval int `json:"unhealthyDrainingInterval,omitempty" yaml:"unhealthyDrainingInterval,omitempty"`
}
