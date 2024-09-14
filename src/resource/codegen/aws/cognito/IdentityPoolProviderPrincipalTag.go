package cognito

type IdentityPoolProviderPrincipalTag struct {
	// An identity pool ID.
	IdentityPoolId string `json:"identityPoolId,omitempty" yaml:"identityPoolId,omitempty"`

	// The name of the identity provider.
	IdentityProviderName string `json:"identityProviderName,omitempty" yaml:"identityProviderName,omitempty"`

	// String to string map of variables.
	PrincipalTags map[string]string `json:"principalTags,omitempty" yaml:"principalTags,omitempty"`

	// use default (username and clientID) attribute mappings.
	UseDefaults bool `json:"useDefaults,omitempty" yaml:"useDefaults,omitempty"`
}
