package types

type Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationTokenSelectionAccessTokenOnly struct {
	// The access token aud claim values that you want to accept in your policy store.
	Audiences []string `json:"audiences,omitempty" yaml:"audiences,omitempty"`

	// The claim that determines the principal in OIDC access tokens.
	PrincipalIdClaim string `json:"principalIdClaim,omitempty" yaml:"principalIdClaim,omitempty"`
}
