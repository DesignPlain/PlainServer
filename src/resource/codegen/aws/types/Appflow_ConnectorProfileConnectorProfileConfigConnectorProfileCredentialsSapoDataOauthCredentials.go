package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOauthCredentials struct {
	//
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	//
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	//
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	//
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOauthCredentialsOauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`
}
