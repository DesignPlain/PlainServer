package types

type Appmesh_getGatewayRouteSpecHttp2Route struct {
	//
	Matches []Appmesh_getGatewayRouteSpecHttp2RouteMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	//
	Actions []Appmesh_getGatewayRouteSpecHttp2RouteAction `json:"actions,omitempty" yaml:"actions,omitempty"`
}
