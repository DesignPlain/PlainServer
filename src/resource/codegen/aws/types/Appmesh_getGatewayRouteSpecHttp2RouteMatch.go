package types

type Appmesh_getGatewayRouteSpecHttp2RouteMatch struct {
	//
	Hostnames []Appmesh_getGatewayRouteSpecHttp2RouteMatchHostname `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`

	//
	Paths []Appmesh_getGatewayRouteSpecHttp2RouteMatchPath `json:"paths,omitempty" yaml:"paths,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	QueryParameters []Appmesh_getGatewayRouteSpecHttp2RouteMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	//
	Headers []Appmesh_getGatewayRouteSpecHttp2RouteMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`
}
