package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2 struct {
	//
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	//
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	//
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	//
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`
}
