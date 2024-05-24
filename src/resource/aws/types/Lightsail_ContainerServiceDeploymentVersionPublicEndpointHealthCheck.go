package types

type Lightsail_ContainerServiceDeploymentVersionPublicEndpointHealthCheck struct {
	// The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. Defaults to "200-499".
	SuccessCodes string `json:"successCodes,omitempty" yaml:"successCodes,omitempty"`

	// The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. Defaults to 2.
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`

	// The number of consecutive health checks failures required before moving the container to the Unhealthy state. Defaults to 2.
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	// The number of consecutive health checks successes required before moving the container to the Healthy state. Defaults to 2.
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	// The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. Defaults to 5.
	IntervalSeconds int `json:"intervalSeconds,omitempty" yaml:"intervalSeconds,omitempty"`

	// The path on the container on which to perform the health check. Defaults to "/".
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
