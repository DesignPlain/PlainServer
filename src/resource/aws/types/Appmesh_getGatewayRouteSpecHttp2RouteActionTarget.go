package types

type Appmesh_getGatewayRouteSpecHttp2RouteActionTarget struct {
	//
	VirtualServices []Appmesh_getGatewayRouteSpecHttp2RouteActionTargetVirtualService `json:"virtualServices,omitempty" yaml:"virtualServices,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
