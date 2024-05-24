package types

type Appmesh_getGatewayRouteSpecGrpcRouteActionTarget struct {
	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	VirtualServices []Appmesh_getGatewayRouteSpecGrpcRouteActionTargetVirtualService `json:"virtualServices,omitempty" yaml:"virtualServices,omitempty"`
}
