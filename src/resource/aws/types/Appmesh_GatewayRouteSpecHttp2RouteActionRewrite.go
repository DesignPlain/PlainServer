package types

type Appmesh_GatewayRouteSpecHttp2RouteActionRewrite struct {
	// Host name to rewrite.
	Hostname Appmesh_GatewayRouteSpecHttp2RouteActionRewriteHostname `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	// Exact path to rewrite.
	Path Appmesh_GatewayRouteSpecHttp2RouteActionRewritePath `json:"path,omitempty" yaml:"path,omitempty"`

	// Specified beginning characters to rewrite.
	Prefix Appmesh_GatewayRouteSpecHttp2RouteActionRewritePrefix `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
