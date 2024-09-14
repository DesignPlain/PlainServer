package compute

type RegionSslCertificate struct {
	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.

	   These are in the same namespace as the managed SSL certificates.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates a unique name beginning with the
	   specified prefix. Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   The write-only private key in PEM format.
	   --Note--: This property is sensitive and will not be displayed in the plan.


	   - - -
	*/
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The Region in which the created regional ssl certificate should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The certificate in PEM format.
	   The certificate chain must be no greater than 5 certs long.
	   The chain must include at least one intermediate cert.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
