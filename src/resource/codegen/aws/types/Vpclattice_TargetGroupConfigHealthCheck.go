package types

type Vpclattice_TargetGroupConfigHealthCheck struct {
	//
	HealthyThresholdCount int `json:"healthyThresholdCount,omitempty" yaml:"healthyThresholdCount,omitempty"`

	// The port used when performing health checks on targets. The default setting is the port that a target receives traffic on.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The protocol used when performing health checks on targets. The possible protocols are `HTTP` and `HTTPS`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The protocol version used when performing health checks on targets. The possible protocol versions are `HTTP1` and `HTTP2`. The default is `HTTP1`.
	ProtocolVersion string `json:"protocolVersion,omitempty" yaml:"protocolVersion,omitempty"`

	// The number of consecutive failed health checks required before considering a target unhealthy. The range is 2–10. The default is 2.
	UnhealthyThresholdCount int `json:"unhealthyThresholdCount,omitempty" yaml:"unhealthyThresholdCount,omitempty"`

	// Indicates whether health checking is enabled. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The approximate amount of time, in seconds, between health checks of an individual target. The range is 5–300 seconds. The default is 30 seconds.
	HealthCheckIntervalSeconds int `json:"healthCheckIntervalSeconds,omitempty" yaml:"healthCheckIntervalSeconds,omitempty"`

	/*
	   The amount of time, in seconds, to wait before reporting a target as unhealthy. The range is 1–120 seconds. The default is 5 seconds.
	   - `healthy_threshold_count ` - (Optional) The number of consecutive successful health checks required before considering an unhealthy target healthy. The range is 2–10. The default is 5.
	*/
	HealthCheckTimeoutSeconds int `json:"healthCheckTimeoutSeconds,omitempty" yaml:"healthCheckTimeoutSeconds,omitempty"`

	// The codes to use when checking for a successful response from a target. These are called _Success codes_ in the console.
	Matcher Vpclattice_TargetGroupConfigHealthCheckMatcher `json:"matcher,omitempty" yaml:"matcher,omitempty"`

	// The destination for health checks on the targets. If the protocol version is HTTP/1.1 or HTTP/2, specify a valid URI (for example, /path?query). The default path is `/`. Health checks are not supported if the protocol version is gRPC, however, you can choose HTTP/1.1 or HTTP/2 and specify a valid URI.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
