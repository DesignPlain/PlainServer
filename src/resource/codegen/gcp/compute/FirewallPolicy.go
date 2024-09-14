package compute

type FirewallPolicy struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The parent of the firewall policy.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.



	   - - -
	*/
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`
}
