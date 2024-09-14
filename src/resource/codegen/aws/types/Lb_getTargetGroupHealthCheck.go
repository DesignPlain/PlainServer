package types

type Lb_getTargetGroupHealthCheck struct {
	//
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	//
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	//
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	//
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	Matcher string `json:"matcher,omitempty" yaml:"matcher,omitempty"`
}
