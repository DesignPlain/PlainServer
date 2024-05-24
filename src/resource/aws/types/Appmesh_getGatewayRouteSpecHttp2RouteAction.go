package types

type Appmesh_getGatewayRouteSpecHttp2RouteAction struct {
	//
	Rewrites []Appmesh_getGatewayRouteSpecHttp2RouteActionRewrite `json:"rewrites,omitempty" yaml:"rewrites,omitempty"`

	//
	Targets []Appmesh_getGatewayRouteSpecHttp2RouteActionTarget `json:"targets,omitempty" yaml:"targets,omitempty"`
}
