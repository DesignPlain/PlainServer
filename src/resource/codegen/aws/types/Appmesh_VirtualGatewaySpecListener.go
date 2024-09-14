package types

type Appmesh_VirtualGatewaySpecListener struct {
	// Port mapping information for the listener.
	PortMapping Appmesh_VirtualGatewaySpecListenerPortMapping `json:"portMapping,omitempty" yaml:"portMapping,omitempty"`

	// Transport Layer Security (TLS) properties for the listener
	Tls Appmesh_VirtualGatewaySpecListenerTls `json:"tls,omitempty" yaml:"tls,omitempty"`

	// Connection pool information for the listener.
	ConnectionPool Appmesh_VirtualGatewaySpecListenerConnectionPool `json:"connectionPool,omitempty" yaml:"connectionPool,omitempty"`

	// Health check information for the listener.
	HealthCheck Appmesh_VirtualGatewaySpecListenerHealthCheck `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`
}
