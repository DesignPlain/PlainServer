package types

type Networkservices_EdgeCacheServiceRoutingPathMatcher struct {
	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name to which this PathMatcher is referred by the HostRule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The routeRules to match against. routeRules support advanced routing behaviour, and can match on paths, headers and query parameters, as well as status codes and HTTP methods.
	   Structure is documented below.
	*/
	RouteRules []Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRule `json:"routeRules,omitempty" yaml:"routeRules,omitempty"`
}
