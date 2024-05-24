package types

type Appmesh_GatewayRouteSpecHttp2RouteMatch struct {
	// Client request query parameters to match on.
	QueryParameters []Appmesh_GatewayRouteSpecHttp2RouteMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	// Client request headers to match on.
	Headers []Appmesh_GatewayRouteSpecHttp2RouteMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Host name to match on.
	Hostname Appmesh_GatewayRouteSpecHttp2RouteMatchHostname `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	// Client request path to match on.
	Path Appmesh_GatewayRouteSpecHttp2RouteMatchPath `json:"path,omitempty" yaml:"path,omitempty"`

	// The port number to match from the request.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Header value sent by the client must begin with the specified characters.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
