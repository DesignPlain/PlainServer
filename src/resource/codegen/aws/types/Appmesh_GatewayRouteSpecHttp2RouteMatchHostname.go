package types

type Appmesh_GatewayRouteSpecHttp2RouteMatchHostname struct {
	// Exact host name to match on.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	// Specified ending characters of the host name to match on.
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`
}
