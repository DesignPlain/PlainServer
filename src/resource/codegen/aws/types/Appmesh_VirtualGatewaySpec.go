package types

type Appmesh_VirtualGatewaySpec struct {
	// Inbound and outbound access logging information for the virtual gateway.
	Logging Appmesh_VirtualGatewaySpecLogging `json:"logging,omitempty" yaml:"logging,omitempty"`

	// Defaults for backends.
	BackendDefaults Appmesh_VirtualGatewaySpecBackendDefaults `json:"backendDefaults,omitempty" yaml:"backendDefaults,omitempty"`

	// Listeners that the mesh endpoint is expected to receive inbound traffic from. You can specify one listener.
	Listeners []Appmesh_VirtualGatewaySpecListener `json:"listeners,omitempty" yaml:"listeners,omitempty"`
}
