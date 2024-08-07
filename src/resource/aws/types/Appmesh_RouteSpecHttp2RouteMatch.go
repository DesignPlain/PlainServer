package types

type Appmesh_RouteSpecHttp2RouteMatch struct {
	// The port number to match from the request.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Header value sent by the client must begin with the specified characters.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Client request query parameters to match on.
	QueryParameters []Appmesh_RouteSpecHttp2RouteMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	// Client request header scheme to match on. Valid values: `http`, `https`.
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`

	// Client request headers to match on.
	Headers []Appmesh_RouteSpecHttp2RouteMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Client request header method to match on. Valid values: `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`.
	Method string `json:"method,omitempty" yaml:"method,omitempty"`

	// Client request path to match on.
	Path Appmesh_RouteSpecHttp2RouteMatchPath `json:"path,omitempty" yaml:"path,omitempty"`
}
