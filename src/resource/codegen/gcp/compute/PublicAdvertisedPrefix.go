package compute

type PublicAdvertisedPrefix struct {
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The IPv4 address to be used for reverse DNS verification.
	DnsVerificationIp string `json:"dnsVerificationIp,omitempty" yaml:"dnsVerificationIp,omitempty"`

	/*
	   The IPv4 address range, in CIDR format, represented by this public advertised prefix.


	   - - -
	*/
	IpCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`

	/*
	   Name of the resource. The name must be 1-63 characters long, and
	   comply with RFC1035. Specifically, the name must be 1-63 characters
	   long and match the regular expression `a-z?`
	   which means the first character must be a lowercase letter, and all
	   following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
