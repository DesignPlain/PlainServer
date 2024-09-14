package types

type Dataproc_ClusterClusterConfigEndpointConfig struct {
	/*
	   The flag to enable http access to specific ports
	   on the cluster from external sources (aka Component Gateway). Defaults to false.
	*/
	EnableHttpPortAccess bool `json:"enableHttpPortAccess,omitempty" yaml:"enableHttpPortAccess,omitempty"`

	// The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true.
	HttpPorts map[string]string `json:"httpPorts,omitempty" yaml:"httpPorts,omitempty"`
}
