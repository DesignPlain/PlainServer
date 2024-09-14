package compute

import types "libds/gcp/types"

type RegionUrlMap struct {
	/*
	   When none of the specified hostRules match, the request is redirected to a URL specified
	   by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	   defaultRouteAction must not be set.
	   Structure is documented below.
	*/
	DefaultUrlRedirect types.Compute_RegionUrlMapDefaultUrlRedirect `json:"defaultUrlRedirect,omitempty" yaml:"defaultUrlRedirect,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The list of HostRules to use against the URL.
	   Structure is documented below.
	*/
	HostRules []types.Compute_RegionUrlMapHostRule `json:"hostRules,omitempty" yaml:"hostRules,omitempty"`

	/*
	   The list of expected URL mappings. Requests to update this UrlMap will
	   succeed only if all of the test cases pass.
	   Structure is documented below.
	*/
	Tests []types.Compute_RegionUrlMapTest `json:"tests,omitempty" yaml:"tests,omitempty"`

	/*
	   defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	   Only one of defaultRouteAction or defaultUrlRedirect must be set.
	   URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction.
	   defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	   Structure is documented below.
	*/
	DefaultRouteAction types.Compute_RegionUrlMapDefaultRouteAction `json:"defaultRouteAction,omitempty" yaml:"defaultRouteAction,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The list of named PathMatchers to use against the URL.
	   Structure is documented below.
	*/
	PathMatchers []types.Compute_RegionUrlMapPathMatcher `json:"pathMatchers,omitempty" yaml:"pathMatchers,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The Region in which the url map should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The full or partial URL of the defaultService resource to which traffic is directed if
	   none of the hostRules match. If defaultRouteAction is additionally specified, advanced
	   routing actions like URL Rewrites, etc. take effect prior to sending the request to the
	   backend. However, if defaultService is specified, defaultRouteAction cannot contain any
	   weightedBackendServices. Conversely, if routeAction specifies any
	   weightedBackendServices, service must not be specified.  Only one of defaultService,
	   defaultUrlRedirect or defaultRouteAction.weightedBackendService must be set.
	*/
	DefaultService string `json:"defaultService,omitempty" yaml:"defaultService,omitempty"`
}
