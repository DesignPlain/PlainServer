package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce struct {
	// The refresh token used to refresh an expired access token.
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`

	// The access token used to access the connector on your behalf.
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// The secret manager ARN, which contains the client ID and client secret of the connected app.
	ClientCredentialsArn string `json:"clientCredentialsArn,omitempty" yaml:"clientCredentialsArn,omitempty"`

	// A JSON web token (JWT) that authorizes access to Salesforce records.
	JwtToken string `json:"jwtToken,omitempty" yaml:"jwtToken,omitempty"`

	// The OAuth 2.0 grant type used by connector for OAuth 2.0 authentication. One of: `AUTHORIZATION_CODE`, `CLIENT_CREDENTIALS`.
	Oauth2GrantType string `json:"oauth2GrantType,omitempty" yaml:"oauth2GrantType,omitempty"`

	// Used by select connectors for which the OAuth workflow is supported. See OAuth Request for more details.
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`
}
