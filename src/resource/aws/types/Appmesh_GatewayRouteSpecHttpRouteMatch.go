package types

type Appmesh_GatewayRouteSpecHttpRouteMatch struct {
	// Client request query parameters to match on.
	QueryParameters []Appmesh_GatewayRouteSpecHttpRouteMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	// Client request headers to match on.
	Headers []Appmesh_GatewayRouteSpecHttpRouteMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Host name to match on.
	Hostname Appmesh_GatewayRouteSpecHttpRouteMatchHostname `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	// Client request path to match on.
	Path Appmesh_GatewayRouteSpecHttpRouteMatchPath `json:"path,omitempty" yaml:"path,omitempty"`

	// The port number to match from the request.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Header value sent by the client must begin with the specified characters.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
