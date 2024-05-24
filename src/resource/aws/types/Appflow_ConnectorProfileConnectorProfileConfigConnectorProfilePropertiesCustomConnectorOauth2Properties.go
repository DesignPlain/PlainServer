package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOauth2Properties struct {
	// Associates your token URL with a map of properties that you define. Use this parameter to provide any additional details that the connector requires to authenticate your request.
	TokenUrlCustomProperties map[string]string `json:"tokenUrlCustomProperties,omitempty" yaml:"tokenUrlCustomProperties,omitempty"`

	// The OAuth 2.0 grant type used by connector for OAuth 2.0 authentication. One of: `AUTHORIZATION_CODE`, `CLIENT_CREDENTIALS`.
	Oauth2GrantType string `json:"oauth2GrantType,omitempty" yaml:"oauth2GrantType,omitempty"`

	// The token URL required for OAuth 2.0 authentication.
	TokenUrl string `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
}
