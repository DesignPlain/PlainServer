package types

type Apprunner_ServiceHealthCheckConfiguration struct {
	// Number of consecutive checks that must succeed before App Runner decides that the service is healthy. Defaults to 1. Minimum value of 1. Maximum value of 20.
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	// Time interval, in seconds, between health checks. Defaults to 5. Minimum value of 1. Maximum value of 20.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	// URL to send requests to for health checks. Defaults to `/`. Minimum length of 0. Maximum length of 51200.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// IP protocol that App Runner uses to perform health checks for your service. Valid values: `TCP`, `HTTP`. Defaults to `TCP`. If you set protocol to `HTTP`, App Runner sends health check requests to the HTTP path specified by `path`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Time, in seconds, to wait for a health check response before deciding it failed. Defaults to 2. Minimum value of  1. Maximum value of 20.
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Number of consecutive checks that must fail before App Runner decides that the service is unhealthy. Defaults to 5. Minimum value of  1. Maximum value of 20.
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`
}
