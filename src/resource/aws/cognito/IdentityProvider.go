package cognito

type IdentityProvider struct {
	// The map of attribute mapping of user pool attributes. [AttributeMapping in AWS API documentation](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-AttributeMapping)
	AttributeMapping map[string]string `json:"attributeMapping,omitempty" yaml:"attributeMapping,omitempty"`

	// The list of identity providers.
	IdpIdentifiers []string `json:"idpIdentifiers,omitempty" yaml:"idpIdentifiers,omitempty"`

	// The map of identity details, such as access token
	ProviderDetails map[string]string `json:"providerDetails,omitempty" yaml:"providerDetails,omitempty"`

	// The provider name
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	// The provider type.  [See AWS API for valid values](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-ProviderType)
	ProviderType string `json:"providerType,omitempty" yaml:"providerType,omitempty"`

	// The user pool id
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`
}
