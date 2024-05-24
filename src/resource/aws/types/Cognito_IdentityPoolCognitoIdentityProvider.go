package types

type Cognito_IdentityPoolCognitoIdentityProvider struct {
	// The provider name for an Amazon Cognito Identity User Pool.
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	// Whether server-side token validation is enabled for the identity provider’s token or not.
	ServerSideTokenCheck bool `json:"serverSideTokenCheck,omitempty" yaml:"serverSideTokenCheck,omitempty"`

	// The client ID for the Amazon Cognito Identity User Pool.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`
}
