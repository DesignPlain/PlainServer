package types

type Appmesh_getRouteSpecHttp2RouteMatch struct {
	//
	Headers []Appmesh_getRouteSpecHttp2RouteMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	//
	Method string `json:"method,omitempty" yaml:"method,omitempty"`

	//
	Paths []Appmesh_getRouteSpecHttp2RouteMatchPath `json:"paths,omitempty" yaml:"paths,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	QueryParameters []Appmesh_getRouteSpecHttp2RouteMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	//
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}
