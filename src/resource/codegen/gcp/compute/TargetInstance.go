package compute

type TargetInstance struct {
	/*
	   The Compute instance VM handling traffic for this target instance.
	   Accepts the instance self-link, relative path
	   (e.g. `projects/project/zones/zone/instances/instance`) or name. If
	   name is given, the zone will default to the given zone or
	   the provider-default zone and the project will default to the
	   provider-level project.


	   - - -
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

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
	   NAT option controlling how IPs are NAT'ed to the instance.
	   Currently only NO_NAT (default value) is supported.
	   Default value is `NO_NAT`.
	   Possible values are: `NO_NAT`.
	*/
	NatPolicy string `json:"natPolicy,omitempty" yaml:"natPolicy,omitempty"`

	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The resource URL for the security policy associated with this target instance.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	// URL of the zone where the target instance resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
