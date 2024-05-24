package types

type Appmesh_getGatewayRouteSpecHttpRouteActionRewrite struct {
	//
	Paths []Appmesh_getGatewayRouteSpecHttpRouteActionRewritePath `json:"paths,omitempty" yaml:"paths,omitempty"`

	//
	Prefixes []Appmesh_getGatewayRouteSpecHttpRouteActionRewritePrefix `json:"prefixes,omitempty" yaml:"prefixes,omitempty"`

	//
	Hostnames []Appmesh_getGatewayRouteSpecHttpRouteActionRewriteHostname `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`
}
