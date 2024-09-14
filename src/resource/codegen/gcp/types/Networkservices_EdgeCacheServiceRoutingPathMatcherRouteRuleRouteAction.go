package types

type Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction struct {
	/*
	   The policy to use for defining caching and signed request behaviour for requests that match this route.
	   Structure is documented below.
	*/
	CdnPolicy Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicy `json:"cdnPolicy,omitempty" yaml:"cdnPolicy,omitempty"`

	/*
	   CORSPolicy defines Cross-Origin-Resource-Sharing configuration, including which CORS response headers will be set.
	   Structure is documented below.
	*/
	CorsPolicy Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCorsPolicy `json:"corsPolicy,omitempty" yaml:"corsPolicy,omitempty"`

	/*
	   The URL rewrite configuration for requests that match this route.
	   Structure is documented below.
	*/
	UrlRewrite Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionUrlRewrite `json:"urlRewrite,omitempty" yaml:"urlRewrite,omitempty"`
}
