package types

type Appmesh_getGatewayRouteSpecHttp2RouteMatchHeader struct {
	//
	Matches []Appmesh_getGatewayRouteSpecHttp2RouteMatchHeaderMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the gateway route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Invert bool `json:"invert,omitempty" yaml:"invert,omitempty"`
}
