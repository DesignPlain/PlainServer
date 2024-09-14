package networksecurity

type TlsInspectionPolicy struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A CA pool resource used to issue interception certificates.
	CaPool string `json:"caPool,omitempty" yaml:"caPool,omitempty"`

	// Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCaSet bool `json:"excludePublicCaSet,omitempty" yaml:"excludePublicCaSet,omitempty"`

	// The location of the tls inspection policy.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Short name of the TlsInspectionPolicy resource to be created.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
