package types

type Compute_URLMapPathMatcherRouteRuleRouteActionUrlRewrite struct {
	/*
	   Prior to forwarding the request to the selected service, the request's host header is replaced
	   with contents of hostRewrite.
	   The value must be between 1 and 255 characters.
	*/
	HostRewrite string `json:"hostRewrite,omitempty" yaml:"hostRewrite,omitempty"`

	/*
	   Prior to forwarding the request to the selected backend service, the matching portion of the
	   request's path is replaced by pathPrefixRewrite.
	   The value must be between 1 and 1024 characters.
	*/
	PathPrefixRewrite string `json:"pathPrefixRewrite,omitempty" yaml:"pathPrefixRewrite,omitempty"`

	/*
	   Prior to forwarding the request to the selected origin, if the
	   request matched a pathTemplateMatch, the matching portion of the
	   request's path is replaced re-written using the pattern specified
	   by pathTemplateRewrite.
	   pathTemplateRewrite must be between 1 and 255 characters
	   (inclusive), must start with a '/', and must only use variables
	   captured by the route's pathTemplate matchers.
	   pathTemplateRewrite may only be used when all of a route's
	   MatchRules specify pathTemplate.
	   Only one of pathPrefixRewrite and pathTemplateRewrite may be
	   specified.
	*/
	PathTemplateRewrite string `json:"pathTemplateRewrite,omitempty" yaml:"pathTemplateRewrite,omitempty"`
}
