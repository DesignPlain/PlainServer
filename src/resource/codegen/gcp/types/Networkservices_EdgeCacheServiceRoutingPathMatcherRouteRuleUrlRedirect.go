package types

type Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect struct {
	/*
	   If set to true, the URL scheme in the redirected request is set to https. If set to false, the URL scheme of the redirected request will remain the same as that of the request.
	   This can only be set if there is at least one (1) edgeSslCertificate set on the service.
	*/
	HttpsRedirect bool `json:"httpsRedirect,omitempty" yaml:"httpsRedirect,omitempty"`

	/*
	   The path that will be used in the redirect response instead of the one that was supplied in the request.
	   pathRedirect cannot be supplied together with prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.
	   The path value must be between 1 and 1024 characters.
	*/
	PathRedirect string `json:"pathRedirect,omitempty" yaml:"pathRedirect,omitempty"`

	/*
	   The prefix that replaces the prefixMatch specified in the routeRule, retaining the remaining portion of the URL before redirecting the request.
	   prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.
	*/
	PrefixRedirect string `json:"prefixRedirect,omitempty" yaml:"prefixRedirect,omitempty"`

	/*
	   The HTTP Status code to use for this RedirectAction.
	   The supported values are:
	*/
	RedirectResponseCode string `json:"redirectResponseCode,omitempty" yaml:"redirectResponseCode,omitempty"`

	/*
	   If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request. If set to false, the query portion of the original URL is retained.

	   - - -
	*/
	StripQuery bool `json:"stripQuery,omitempty" yaml:"stripQuery,omitempty"`

	// The host that will be used in the redirect response instead of the one that was supplied in the request.
	HostRedirect string `json:"hostRedirect,omitempty" yaml:"hostRedirect,omitempty"`
}
