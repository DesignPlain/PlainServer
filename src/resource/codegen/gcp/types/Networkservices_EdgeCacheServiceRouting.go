package types

type Networkservices_EdgeCacheServiceRouting struct {
	/*
	   The list of hostRules to match against. These rules define which hostnames the EdgeCacheService will match against, and which route configurations apply.
	   Structure is documented below.
	*/
	HostRules []Networkservices_EdgeCacheServiceRoutingHostRule `json:"hostRules,omitempty" yaml:"hostRules,omitempty"`

	/*
	   The list of pathMatchers referenced via name by hostRules. PathMatcher is used to match the path portion of the URL when a HostRule matches the URL's host portion.
	   Structure is documented below.
	*/
	PathMatchers []Networkservices_EdgeCacheServiceRoutingPathMatcher `json:"pathMatchers,omitempty" yaml:"pathMatchers,omitempty"`
}
