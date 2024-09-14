package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigEndpointConfig struct {
	// If true, enable http access to specific ports on the cluster from external sources. Defaults to false.
	EnableHttpPortAccess bool `json:"enableHttpPortAccess,omitempty" yaml:"enableHttpPortAccess,omitempty"`

	// Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true.
	HttpPorts map[string]string `json:"httpPorts,omitempty" yaml:"httpPorts,omitempty"`
}
