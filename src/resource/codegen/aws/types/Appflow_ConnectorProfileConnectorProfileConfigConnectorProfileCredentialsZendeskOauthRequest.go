package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequest struct {
	// The code provided by the connector when it has been authenticated via the connected app.
	AuthCode string `json:"authCode,omitempty" yaml:"authCode,omitempty"`

	// The URL to which the authentication server redirects the browser after authorization has been granted.
	RedirectUri string `json:"redirectUri,omitempty" yaml:"redirectUri,omitempty"`
}
