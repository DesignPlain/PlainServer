package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode struct {
	// The access token used to access the connector on your behalf.
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// Used by select connectors for which the OAuth workflow is supported. See OAuth Request for more details.
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycodeOauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`

	// The refresh token used to refresh an expired access token.
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`
}
