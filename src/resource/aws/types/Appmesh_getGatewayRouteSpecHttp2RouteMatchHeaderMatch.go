package types

type Appmesh_getGatewayRouteSpecHttp2RouteMatchHeaderMatch struct {
	//
	Ranges []Appmesh_getGatewayRouteSpecHttp2RouteMatchHeaderMatchRange `json:"ranges,omitempty" yaml:"ranges,omitempty"`

	//
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	//
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`

	//
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
