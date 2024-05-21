package types

type Cloudrunv2_getServiceTemplateContainerLivenessProbe struct {
	// Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. Maximum value is 3600. Must be smaller than periodSeconds. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`

	// Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
	FailureThreshold int `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`

	// GRPC specifies an action involving a GRPC port.
	Grpcs []Cloudrunv2_getServiceTemplateContainerLivenessProbeGrpc `json:"grpcs,omitempty" yaml:"grpcs,omitempty"`

	// HTTPGet specifies the http request to perform.
	HttpGets []Cloudrunv2_getServiceTemplateContainerLivenessProbeHttpGet `json:"httpGets,omitempty" yaml:"httpGets,omitempty"`

	// Number of seconds after the container has started before the probe is initiated. Defaults to 0 seconds. Minimum value is 0. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	InitialDelaySeconds int `json:"initialDelaySeconds,omitempty" yaml:"initialDelaySeconds,omitempty"`

	// How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. Must be greater or equal than timeoutSeconds
	PeriodSeconds int `json:"periodSeconds,omitempty" yaml:"periodSeconds,omitempty"`

	// TCPSocketAction describes an action based on opening a socket
	TcpSockets []Cloudrunv2_getServiceTemplateContainerLivenessProbeTcpSocket `json:"tcpSockets,omitempty" yaml:"tcpSockets,omitempty"`
}
