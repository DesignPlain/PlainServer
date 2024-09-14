package types

type Appmesh_getGatewayRouteSpecHttpRouteMatchHeader struct {
	//
	Invert bool `json:"invert,omitempty" yaml:"invert,omitempty"`

	//
	Matches []Appmesh_getGatewayRouteSpecHttpRouteMatchHeaderMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the gateway route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
