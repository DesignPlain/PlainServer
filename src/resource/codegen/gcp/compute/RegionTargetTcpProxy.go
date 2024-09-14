package compute

type RegionTargetTcpProxy struct {
	/*
	   The Region in which the created target TCP proxy should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   A reference to the BackendService resource.


	   - - -
	*/
	BackendService string `json:"backendService,omitempty" yaml:"backendService,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

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

	/*
	   Specifies the type of proxy header to append before sending data to
	   the backend.
	   Default value is `NONE`.
	   Possible values are: `NONE`, `PROXY_V1`.
	*/
	ProxyHeader string `json:"proxyHeader,omitempty" yaml:"proxyHeader,omitempty"`
}
