package types

type Elb_LoadBalancerHealthCheck struct {
	// The number of checks before the instance is declared healthy.
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	// The interval between checks.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	/*
	   The target of the check. Valid pattern is "${PROTOCOL}:${PORT}${PATH}", where PROTOCOL
	   values are:
	   - `HTTP`, `HTTPS` - PORT and PATH are required
	   - `TCP`, `SSL` - PORT is required, PATH is not supported
	*/
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// The length of time before the check times out.
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// The number of checks before the instance is declared unhealthy.
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`
}
