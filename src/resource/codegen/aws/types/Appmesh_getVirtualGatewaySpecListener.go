package types

type Appmesh_getVirtualGatewaySpecListener struct {
	//
	ConnectionPools []Appmesh_getVirtualGatewaySpecListenerConnectionPool `json:"connectionPools,omitempty" yaml:"connectionPools,omitempty"`

	//
	HealthChecks []Appmesh_getVirtualGatewaySpecListenerHealthCheck `json:"healthChecks,omitempty" yaml:"healthChecks,omitempty"`

	//
	PortMappings []Appmesh_getVirtualGatewaySpecListenerPortMapping `json:"portMappings,omitempty" yaml:"portMappings,omitempty"`

	//
	Tls []Appmesh_getVirtualGatewaySpecListenerTl `json:"tls,omitempty" yaml:"tls,omitempty"`
}
