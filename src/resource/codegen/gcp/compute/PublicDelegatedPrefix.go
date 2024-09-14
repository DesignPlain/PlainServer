package compute

type PublicDelegatedPrefix struct {
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The IPv4 address range, in CIDR format, represented by this public advertised prefix.


	   - - -
	*/
	IpCidrRange string `json:"ipCidrRange,omitempty" yaml:"ipCidrRange,omitempty"`

	// If true, the prefix will be live migrated.
	IsLiveMigration bool `json:"isLiveMigration,omitempty" yaml:"isLiveMigration,omitempty"`

	/*
	   Name of the resource. The name must be 1-63 characters long, and
	   comply with RFC1035. Specifically, the name must be 1-63 characters
	   long and match the regular expression `a-z?`
	   which means the first character must be a lowercase letter, and all
	   following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix string `json:"parentPrefix,omitempty" yaml:"parentPrefix,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A region where the prefix will reside.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
