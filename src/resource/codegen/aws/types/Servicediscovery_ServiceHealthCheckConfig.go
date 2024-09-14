package types

type Servicediscovery_ServiceHealthCheckConfig struct {
	// The number of consecutive health checks. Maximum value of 10.
	FailureThreshold int `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`

	// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
	ResourcePath string `json:"resourcePath,omitempty" yaml:"resourcePath,omitempty"`

	// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
