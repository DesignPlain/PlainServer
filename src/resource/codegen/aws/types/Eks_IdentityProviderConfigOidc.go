package types

type Eks_IdentityProviderConfigOidc struct {
	// A prefix that is prepended to username claims.
	UsernamePrefix string `json:"usernamePrefix,omitempty" yaml:"usernamePrefix,omitempty"`

	// Client ID for the OpenID Connect identity provider.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The JWT claim that the provider will use to return groups.
	GroupsClaim string `json:"groupsClaim,omitempty" yaml:"groupsClaim,omitempty"`

	// A prefix that is prepended to group claims e.g., `oidc:`.
	GroupsPrefix string `json:"groupsPrefix,omitempty" yaml:"groupsPrefix,omitempty"`

	// The name of the identity provider config.
	IdentityProviderConfigName string `json:"identityProviderConfigName,omitempty" yaml:"identityProviderConfigName,omitempty"`

	// Issuer URL for the OpenID Connect identity provider.
	IssuerUrl string `json:"issuerUrl,omitempty" yaml:"issuerUrl,omitempty"`

	// The key value pairs that describe required claims in the identity token.
	RequiredClaims map[string]string `json:"requiredClaims,omitempty" yaml:"requiredClaims,omitempty"`

	// The JWT claim that the provider will use as the username.
	UsernameClaim string `json:"usernameClaim,omitempty" yaml:"usernameClaim,omitempty"`
}
