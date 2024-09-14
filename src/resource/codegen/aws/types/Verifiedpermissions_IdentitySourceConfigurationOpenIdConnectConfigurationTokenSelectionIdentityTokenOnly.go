package types

type Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationTokenSelectionIdentityTokenOnly struct {
	// The claim that determines the principal in OIDC access tokens.
	PrincipalIdClaim string `json:"principalIdClaim,omitempty" yaml:"principalIdClaim,omitempty"`

	// The ID token audience, or client ID, claim values that you want to accept in your policy store from an OIDC identity provider.
	ClientIds []string `json:"clientIds,omitempty" yaml:"clientIds,omitempty"`
}
