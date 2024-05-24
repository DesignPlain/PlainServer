package types

type Cloudwatch_EventEndpointRoutingConfigFailoverConfigPrimary struct {
	// The ARN of the health check used by the endpoint to determine whether failover is triggered.
	HealthCheck string `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`
}
