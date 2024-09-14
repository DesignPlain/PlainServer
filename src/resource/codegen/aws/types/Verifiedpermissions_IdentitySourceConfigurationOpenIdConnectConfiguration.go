package types

type Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfiguration struct {
	// A descriptive string that you want to prefix to user entities from your OIDC identity provider.
	EntityIdPrefix string `json:"entityIdPrefix,omitempty" yaml:"entityIdPrefix,omitempty"`

	// The type of entity that a policy store maps to groups from an Amazon Cognito user pool identity source. See Group Configuration below.
	GroupConfiguration Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationGroupConfiguration `json:"groupConfiguration,omitempty" yaml:"groupConfiguration,omitempty"`

	// The issuer URL of an OIDC identity provider. This URL must have an OIDC discovery endpoint at the path `.well-known/openid-configuration`.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// The token type that you want to process from your OIDC identity provider. Your policy store can process either identity (ID) or access tokens from a given OIDC identity source. See Token Selection below.
	TokenSelection Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationTokenSelection `json:"tokenSelection,omitempty" yaml:"tokenSelection,omitempty"`
}
