package types

type Appmesh_GatewayRouteSpecHttpRouteActionRewrite struct {
	// Host name to rewrite.
	Hostname Appmesh_GatewayRouteSpecHttpRouteActionRewriteHostname `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	// Exact path to rewrite.
	Path Appmesh_GatewayRouteSpecHttpRouteActionRewritePath `json:"path,omitempty" yaml:"path,omitempty"`

	// Specified beginning characters to rewrite.
	Prefix Appmesh_GatewayRouteSpecHttpRouteActionRewritePrefix `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
