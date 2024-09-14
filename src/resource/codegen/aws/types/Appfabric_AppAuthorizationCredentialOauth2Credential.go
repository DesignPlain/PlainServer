package types

type Appfabric_AppAuthorizationCredentialOauth2Credential struct {
	// The client ID of the client application.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The client secret of the client application.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
}
