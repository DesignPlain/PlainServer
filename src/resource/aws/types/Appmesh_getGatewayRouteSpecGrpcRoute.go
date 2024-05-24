package types

type Appmesh_getGatewayRouteSpecGrpcRoute struct {
	//
	Matches []Appmesh_getGatewayRouteSpecGrpcRouteMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	//
	Actions []Appmesh_getGatewayRouteSpecGrpcRouteAction `json:"actions,omitempty" yaml:"actions,omitempty"`
}
