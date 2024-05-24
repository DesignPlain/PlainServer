package types

type Alb_getTargetGroupHealthCheck struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	//
	Matcher string `json:"matcher,omitempty" yaml:"matcher,omitempty"`

	//
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	//
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	//
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`
}
