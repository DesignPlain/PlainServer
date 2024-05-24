package types

type Appmesh_getGatewayRouteSpecHttpRouteMatch struct {
	//
	Hostnames []Appmesh_getGatewayRouteSpecHttpRouteMatchHostname `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`

	//
	Paths []Appmesh_getGatewayRouteSpecHttpRouteMatchPath `json:"paths,omitempty" yaml:"paths,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	QueryParameters []Appmesh_getGatewayRouteSpecHttpRouteMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	//
	Headers []Appmesh_getGatewayRouteSpecHttpRouteMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`
}
