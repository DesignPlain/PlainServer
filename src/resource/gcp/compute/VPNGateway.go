package compute

type VPNGateway struct {
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035.  Specifically, the name must be 1-63 characters long and
	   match the regular expression `a-z?` which means
	   the first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The network this VPN gateway is accepting traffic for.


	   - - -
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region this gateway should sit in.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
