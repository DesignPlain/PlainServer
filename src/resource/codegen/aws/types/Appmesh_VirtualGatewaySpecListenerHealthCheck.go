package types

type Appmesh_VirtualGatewaySpecListenerHealthCheck struct {
	// Time period in milliseconds between each health check execution.
	IntervalMillis int `json:"intervalMillis,omitempty" yaml:"intervalMillis,omitempty"`

	// Destination path for the health check request. This is only required if the specified protocol is `http` or `http2`.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Destination port for the health check request. This port must match the port defined in the `port_mapping` for the listener.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Protocol for the health check request. Valid values are `http`, `http2`, and `grpc`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Amount of time to wait when receiving a response from the health check, in milliseconds.
	TimeoutMillis int `json:"timeoutMillis,omitempty" yaml:"timeoutMillis,omitempty"`

	// Number of consecutive failed health checks that must occur before declaring a virtual gateway unhealthy.
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	// Number of consecutive successful health checks that must occur before declaring listener healthy.
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`
}
