package compute

import types "libds/gcp/types"

type URLMap struct {
	/*
	   When none of the specified hostRules match, the request is redirected to a URL specified
	   by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	   defaultRouteAction must not be set.
	   Structure is documented below.
	*/
	DefaultUrlRedirect types.Compute_URLMapDefaultUrlRedirect `json:"defaultUrlRedirect,omitempty" yaml:"defaultUrlRedirect,omitempty"`

	/*
	   Specifies changes to request and response headers that need to take effect for
	   the selected backendService. The headerAction specified here take effect after
	   headerAction specified under pathMatcher.
	   Structure is documented below.
	*/
	HeaderAction types.Compute_URLMapHeaderAction `json:"headerAction,omitempty" yaml:"headerAction,omitempty"`

	/*
	   The list of named PathMatchers to use against the URL.
	   Structure is documented below.
	*/
	PathMatchers []types.Compute_URLMapPathMatcher `json:"pathMatchers,omitempty" yaml:"pathMatchers,omitempty"`

	// The backend service or backend bucket to use when none of the given rules match.
	DefaultService string `json:"defaultService,omitempty" yaml:"defaultService,omitempty"`

	/*
	   An optional description of this resource. Provide this property when you create
	   the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The list of HostRules to use against the URL.
	   Structure is documented below.
	*/
	HostRules []types.Compute_URLMapHostRule `json:"hostRules,omitempty" yaml:"hostRules,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is created. The
	   name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	   name must be 1-63 characters long and match the regular expression
	   `a-z?` which means the first character must be a lowercase
	   letter, and all following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The list of expected URL mapping tests. Request to update this UrlMap will
	   succeed only if all of the test cases pass. You can specify a maximum of 100
	   tests per UrlMap.
	   Structure is documented below.
	*/
	Tests []types.Compute_URLMapTest `json:"tests,omitempty" yaml:"tests,omitempty"`

	/*
	   defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions
	   like URL rewrites, header transformations, etc. prior to forwarding the request to the selected backend.
	   If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService
	   is set, defaultRouteAction cannot contain any weightedBackendServices.
	   Only one of defaultRouteAction or defaultUrlRedirect must be set.
	   Structure is documented below.
	*/
	DefaultRouteAction types.Compute_URLMapDefaultRouteAction `json:"defaultRouteAction,omitempty" yaml:"defaultRouteAction,omitempty"`
}
