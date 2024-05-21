package compute

type NetworkFirewallPolicy struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   User-provided name of the Network firewall policy. The name should be unique in the project in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
