package types

type Appmesh_getVirtualGatewaySpec struct {
	//
	Listeners []Appmesh_getVirtualGatewaySpecListener `json:"listeners,omitempty" yaml:"listeners,omitempty"`

	//
	Loggings []Appmesh_getVirtualGatewaySpecLogging `json:"loggings,omitempty" yaml:"loggings,omitempty"`

	//
	BackendDefaults []Appmesh_getVirtualGatewaySpecBackendDefault `json:"backendDefaults,omitempty" yaml:"backendDefaults,omitempty"`
}
