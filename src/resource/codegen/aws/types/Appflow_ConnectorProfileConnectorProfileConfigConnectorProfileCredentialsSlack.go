package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack struct {
	//
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	//
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	//
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	//
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`
}