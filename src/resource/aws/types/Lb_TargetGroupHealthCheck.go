package types

type Lb_TargetGroupHealthCheck struct {
	/*
	   The HTTP or gRPC codes to use when checking for a successful response from a target.
	   The `health_check.protocol` must be one of `HTTP` or `HTTPS` or the `target_type` must be `lambda`.
	   Values can be comma-separated individual values (e.g., "200,202") or a range of values (e.g., "200-299").
	   - For gRPC-based target groups (i.e., the `protocol` is one of `HTTP` or `HTTPS` and the `protocol_version` is `GRPC`), values can be between `0` and `99`. The default is `12`.
	   - When used with an Application Load Balancer (i.e., the `protocol` is one of `HTTP` or `HTTPS` and the `protocol_version` is not `GRPC`), values can be between `200` and `499`. The default is `200`.
	   - When used with a Network Load Balancer (i.e., the `protocol` is one of `TCP`, `TCP_UDP`, `UDP`, or `TLS`), values can be between `200` and `599`. The default is `200-399`.
	   - When the `target_type` is `lambda`, values can be between `200` and `499`. The default is `200`.
	*/
	Matcher string `json:"matcher,omitempty" yaml:"matcher,omitempty"`

	/*
	   Protocol the load balancer uses when performing health checks on targets.
	   Must be one of `TCP`, `HTTP`, or `HTTPS`.
	   The `TCP` protocol is not supported for health checks if the protocol of the target group is `HTTP` or `HTTPS`.
	   Default is `HTTP`.
	   Cannot be specified when the `target_type` is `lambda`.
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Whether health checks are enabled. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Approximate amount of time, in seconds, between health checks of an individual target. The range is 5-300. For `lambda` target groups, it needs to be greater than the timeout of the underlying `lambda`. Defaults to 30.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	/*
	   The port the load balancer uses when performing health checks on targets.
	   Valid values are either `traffic-port`, to use the same port as the target group, or a valid port number between `1` and `65536`.
	   Default is `traffic-port`.
	*/
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	// Amount of time, in seconds, during which no response from a target means a failed health check. The range is 2–120 seconds. For target groups with a protocol of HTTP, the default is 6 seconds. For target groups with a protocol of TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a protocol of GENEVE, the default is 5 seconds. If the target type is lambda, the default is 30 seconds.
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Number of consecutive health check failures required before considering a target unhealthy. The range is 2-10. Defaults to 3.
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	// Number of consecutive health check successes required before considering a target healthy. The range is 2-10. Defaults to 3.
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	/*
	   Destination for the health check request. Required for HTTP/HTTPS ALB and HTTP NLB. Only applies to HTTP/HTTPS.
	   - For HTTP and HTTPS health checks, the default is `/`.
	   - For gRPC health checks, the default is `/Amazon Web Services.ALB/healthcheck`.
	*/
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
