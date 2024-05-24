package types

type Appmesh_VirtualNodeSpecListener struct {
	// Connection pool information for the listener.
	ConnectionPool Appmesh_VirtualNodeSpecListenerConnectionPool `json:"connectionPool,omitempty" yaml:"connectionPool,omitempty"`

	// Health check information for the listener.
	HealthCheck Appmesh_VirtualNodeSpecListenerHealthCheck `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	// Outlier detection information for the listener.
	OutlierDetection Appmesh_VirtualNodeSpecListenerOutlierDetection `json:"outlierDetection,omitempty" yaml:"outlierDetection,omitempty"`

	// Port mapping information for the listener.
	PortMapping Appmesh_VirtualNodeSpecListenerPortMapping `json:"portMapping,omitempty" yaml:"portMapping,omitempty"`

	// Timeouts for different protocols.
	Timeout Appmesh_VirtualNodeSpecListenerTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Transport Layer Security (TLS) properties for the listener
	Tls Appmesh_VirtualNodeSpecListenerTls `json:"tls,omitempty" yaml:"tls,omitempty"`
}
