package types

type Appmesh_getGatewayRouteSpecHttp2RouteActionTarget struct {
	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	VirtualServices []Appmesh_getGatewayRouteSpecHttp2RouteActionTargetVirtualService `json:"virtualServices,omitempty" yaml:"virtualServices,omitempty"`
}
