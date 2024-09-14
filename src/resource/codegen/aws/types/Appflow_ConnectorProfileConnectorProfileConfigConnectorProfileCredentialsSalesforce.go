package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce struct {
	// A JSON web token (JWT) that authorizes access to Salesforce records.
	JwtToken string `json:"jwtToken,omitempty" yaml:"jwtToken,omitempty"`

	//
	Oauth2GrantType string `json:"oauth2GrantType,omitempty" yaml:"oauth2GrantType,omitempty"`

	//
	OauthRequest Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequest `json:"oauthRequest,omitempty" yaml:"oauthRequest,omitempty"`

	//
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`

	//
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// The secret manager ARN, which contains the client ID and client secret of the connected app.
	ClientCredentialsArn string `json:"clientCredentialsArn,omitempty" yaml:"clientCredentialsArn,omitempty"`
}
