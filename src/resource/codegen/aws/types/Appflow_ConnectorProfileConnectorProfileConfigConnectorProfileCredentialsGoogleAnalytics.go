package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics struct {
	//
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	//
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	//
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	//
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`
}
