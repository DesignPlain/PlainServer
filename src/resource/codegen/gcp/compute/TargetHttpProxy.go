package compute

type TargetHttpProxy struct {
	/*
	   A reference to the UrlMap resource that defines the mapping from URL
	   to the BackendService.


	   - - -
	*/
	UrlMap string `json:"urlMap,omitempty" yaml:"urlMap,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specifies how long to keep a connection open, after completing a response,
	   while there is no matching traffic (in seconds). If an HTTP keepalive is
	   not specified, a default value (610 seconds) will be used. For Global
	   external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	   the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	   load balancer (classic), this option is not available publicly.
	*/
	HttpKeepAliveTimeoutSec int `json:"httpKeepAliveTimeoutSec,omitempty" yaml:"httpKeepAliveTimeoutSec,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   This field only applies when the forwarding rule that references
	   this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	*/
	ProxyBind bool `json:"proxyBind,omitempty" yaml:"proxyBind,omitempty"`
}
