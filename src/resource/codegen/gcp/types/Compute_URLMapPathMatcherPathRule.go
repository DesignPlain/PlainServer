package types

type Compute_URLMapPathMatcherPathRule struct {
	/*
	   The list of path patterns to match. Each must start with / and the only place a
	   \- is allowed is at the end following a /. The string fed to the path matcher
	   does not include any text after the first ? or #, and those chars are not
	   allowed here.
	*/
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`

	/*
	   In response to a matching path, the load balancer performs advanced routing
	   actions like URL rewrites, header transformations, etc. prior to forwarding the
	   request to the selected backend. If routeAction specifies any
	   weightedBackendServices, service must not be set. Conversely if service is set,
	   routeAction cannot contain any  weightedBackendServices. Only one of routeAction
	   or urlRedirect must be set.
	   Structure is documented below.
	*/
	RouteAction Compute_URLMapPathMatcherPathRuleRouteAction `json:"routeAction,omitempty" yaml:"routeAction,omitempty"`

	// The backend service or backend bucket to use if any of the given paths match.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   When a path pattern is matched, the request is redirected to a URL specified
	   by urlRedirect. If urlRedirect is specified, service or routeAction must not
	   be set.
	   Structure is documented below.
	*/
	UrlRedirect Compute_URLMapPathMatcherPathRuleUrlRedirect `json:"urlRedirect,omitempty" yaml:"urlRedirect,omitempty"`
}
