package types

type Cognito_getIdentityPoolCognitoIdentityProvider struct {
	//
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	//
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	//
	ServerSideTokenCheck bool `json:"serverSideTokenCheck,omitempty" yaml:"serverSideTokenCheck,omitempty"`
}
