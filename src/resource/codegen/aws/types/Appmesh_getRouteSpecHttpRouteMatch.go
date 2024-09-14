package types

type Appmesh_getRouteSpecHttpRouteMatch struct {
	//
	QueryParameters []Appmesh_getRouteSpecHttpRouteMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	//
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`

	//
	Headers []Appmesh_getRouteSpecHttpRouteMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	//
	Method string `json:"method,omitempty" yaml:"method,omitempty"`

	//
	Paths []Appmesh_getRouteSpecHttpRouteMatchPath `json:"paths,omitempty" yaml:"paths,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
