package types

type Elb_getLoadBalancerHealthCheck struct {
	//
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	//
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	//
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	//
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	//
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`
}
