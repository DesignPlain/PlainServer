package types

type Appmesh_getGatewayRouteSpecHttpRouteActionTarget struct {
	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	VirtualServices []Appmesh_getGatewayRouteSpecHttpRouteActionTargetVirtualService `json:"virtualServices,omitempty" yaml:"virtualServices,omitempty"`
}
