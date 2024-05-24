package types

type Appmesh_getGatewayRouteSpecHttp2RouteActionRewrite struct {
	//
	Paths []Appmesh_getGatewayRouteSpecHttp2RouteActionRewritePath `json:"paths,omitempty" yaml:"paths,omitempty"`

	//
	Prefixes []Appmesh_getGatewayRouteSpecHttp2RouteActionRewritePrefix `json:"prefixes,omitempty" yaml:"prefixes,omitempty"`

	//
	Hostnames []Appmesh_getGatewayRouteSpecHttp2RouteActionRewriteHostname `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`
}
