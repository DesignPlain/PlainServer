package types

type Cloudrun_getServiceTemplateSpecContainerStartupProbe struct {
	// TcpSocket specifies an action involving a TCP port.
	TcpSockets []Cloudrun_getServiceTemplateSpecContainerStartupProbeTcpSocket `json:"tcpSockets,omitempty" yaml:"tcpSockets,omitempty"`

	/*
	   Number of seconds after which the probe times out.
	   Defaults to 1 second. Minimum value is 1. Maximum value is 3600.
	   Must be smaller than periodSeconds.
	*/
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`

	/*
	   Minimum consecutive failures for the probe to be considered failed after
	   having succeeded. Defaults to 3. Minimum value is 1.
	*/
	FailureThreshold int `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`

	// GRPC specifies an action involving a GRPC port.
	Grpcs []Cloudrun_getServiceTemplateSpecContainerStartupProbeGrpc `json:"grpcs,omitempty" yaml:"grpcs,omitempty"`

	// HttpGet specifies the http request to perform.
	HttpGets []Cloudrun_getServiceTemplateSpecContainerStartupProbeHttpGet `json:"httpGets,omitempty" yaml:"httpGets,omitempty"`

	/*
	   Number of seconds after the container has started before the probe is
	   initiated.
	   Defaults to 0 seconds. Minimum value is 0. Maximum value is 240.
	*/
	InitialDelaySeconds int `json:"initialDelaySeconds,omitempty" yaml:"initialDelaySeconds,omitempty"`

	/*
	   How often (in seconds) to perform the probe.
	   Default to 10 seconds. Minimum value is 1. Maximum value is 240.
	*/
	PeriodSeconds int `json:"periodSeconds,omitempty" yaml:"periodSeconds,omitempty"`
}
