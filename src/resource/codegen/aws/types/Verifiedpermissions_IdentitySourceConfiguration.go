package types

type Verifiedpermissions_IdentitySourceConfiguration struct {
	// Specifies the configuration details of an Amazon Cognito user pool that Verified Permissions can use as a source of authenticated identities as entities. See Cognito User Pool Configuration below.
	CognitoUserPoolConfiguration Verifiedpermissions_IdentitySourceConfigurationCognitoUserPoolConfiguration `json:"cognitoUserPoolConfiguration,omitempty" yaml:"cognitoUserPoolConfiguration,omitempty"`

	// Specifies the configuration details of an OpenID Connect (OIDC) identity provider, or identity source, that Verified Permissions can use to generate entities from authenticated identities. See Open ID Connect Configuration below.
	OpenIdConnectConfiguration Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfiguration `json:"openIdConnectConfiguration,omitempty" yaml:"openIdConnectConfiguration,omitempty"`
}
