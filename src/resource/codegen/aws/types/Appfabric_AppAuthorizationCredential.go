package types

type Appfabric_AppAuthorizationCredential struct {
	// Contains OAuth2 client credential information.
	Oauth2Credential Appfabric_AppAuthorizationCredentialOauth2Credential `json:"oauth2Credential,omitempty" yaml:"oauth2Credential,omitempty"`

	// Contains API key credential information.
	ApiKeyCredentials []Appfabric_AppAuthorizationCredentialApiKeyCredential `json:"apiKeyCredentials,omitempty" yaml:"apiKeyCredentials,omitempty"`
}
