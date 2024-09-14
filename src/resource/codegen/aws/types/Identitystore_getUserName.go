package types

type Identitystore_getUserName struct {
	// The family name of the user.
	FamilyName string `json:"familyName,omitempty" yaml:"familyName,omitempty"`

	// The name that is typically displayed when the name is shown for display.
	Formatted string `json:"formatted,omitempty" yaml:"formatted,omitempty"`

	// The given name of the user.
	GivenName string `json:"givenName,omitempty" yaml:"givenName,omitempty"`

	// The honorific prefix of the user.
	HonorificPrefix string `json:"honorificPrefix,omitempty" yaml:"honorificPrefix,omitempty"`

	// The honorific suffix of the user.
	HonorificSuffix string `json:"honorificSuffix,omitempty" yaml:"honorificSuffix,omitempty"`

	// The middle name of the user.
	MiddleName string `json:"middleName,omitempty" yaml:"middleName,omitempty"`
}
