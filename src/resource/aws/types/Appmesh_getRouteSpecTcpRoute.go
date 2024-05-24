package types

type Appmesh_getRouteSpecTcpRoute struct {
	//
	Actions []Appmesh_getRouteSpecTcpRouteAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	//
	Matches []Appmesh_getRouteSpecTcpRouteMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	//
	Timeouts []Appmesh_getRouteSpecTcpRouteTimeout `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
