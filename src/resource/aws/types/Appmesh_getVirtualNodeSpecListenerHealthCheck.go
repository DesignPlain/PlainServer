package types

type Appmesh_getVirtualNodeSpecListenerHealthCheck struct {
	//
	TimeoutMillis int `json:"timeoutMillis,omitempty" yaml:"timeoutMillis,omitempty"`

	//
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" yaml:"unhealthyThreshold,omitempty"`

	//
	HealthyThreshold int `json:"healthyThreshold,omitempty" yaml:"healthyThreshold,omitempty"`

	//
	IntervalMillis int `json:"intervalMillis,omitempty" yaml:"intervalMillis,omitempty"`

	//
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
