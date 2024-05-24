package types

type Appmesh_VirtualNodeSpec struct {
	// Defaults for backends.
	BackendDefaults Appmesh_VirtualNodeSpecBackendDefaults `json:"backendDefaults,omitempty" yaml:"backendDefaults,omitempty"`

	// Backends to which the virtual node is expected to send outbound traffic.
	Backends []Appmesh_VirtualNodeSpecBackend `json:"backends,omitempty" yaml:"backends,omitempty"`

	// Listeners from which the virtual node is expected to receive inbound traffic.
	Listeners []Appmesh_VirtualNodeSpecListener `json:"listeners,omitempty" yaml:"listeners,omitempty"`

	// Inbound and outbound access logging information for the virtual node.
	Logging Appmesh_VirtualNodeSpecLogging `json:"logging,omitempty" yaml:"logging,omitempty"`

	// Service discovery information for the virtual node.
	ServiceDiscovery Appmesh_VirtualNodeSpecServiceDiscovery `json:"serviceDiscovery,omitempty" yaml:"serviceDiscovery,omitempty"`
}
