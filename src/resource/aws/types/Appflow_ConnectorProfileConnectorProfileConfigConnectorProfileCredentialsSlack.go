package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack struct {
	// Used by select connectors for which the OAuth workflow is supported. See OAuth Request for more details.
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`

	// The access token used to access the connector on your behalf.
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// The identifier for the desired client.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
}
