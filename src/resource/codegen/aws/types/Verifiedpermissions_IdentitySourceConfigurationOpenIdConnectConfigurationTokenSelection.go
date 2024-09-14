package types

type Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationTokenSelection struct {
	// The OIDC configuration for processing access tokens. See Access Token Only below.
	AccessTokenOnly Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationTokenSelectionAccessTokenOnly `json:"accessTokenOnly,omitempty" yaml:"accessTokenOnly,omitempty"`

	// The OIDC configuration for processing identity (ID) tokens. See Identity Token Only below.
	IdentityTokenOnly Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationTokenSelectionIdentityTokenOnly `json:"identityTokenOnly,omitempty" yaml:"identityTokenOnly,omitempty"`
}
