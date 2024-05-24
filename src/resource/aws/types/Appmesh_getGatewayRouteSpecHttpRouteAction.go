package types

type Appmesh_getGatewayRouteSpecHttpRouteAction struct {
	//
	Rewrites []Appmesh_getGatewayRouteSpecHttpRouteActionRewrite `json:"rewrites,omitempty" yaml:"rewrites,omitempty"`

	//
	Targets []Appmesh_getGatewayRouteSpecHttpRouteActionTarget `json:"targets,omitempty" yaml:"targets,omitempty"`
}
