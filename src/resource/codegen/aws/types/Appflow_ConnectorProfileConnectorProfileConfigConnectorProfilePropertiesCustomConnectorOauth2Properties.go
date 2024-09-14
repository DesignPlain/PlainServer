package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOauth2Properties struct {
	//
	Oauth2GrantType string `json:"oauth2GrantType,omitempty" yaml:"oauth2GrantType,omitempty"`

	//
	TokenUrl string `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`

	// Associates your token URL with a map of properties that you define. Use this parameter to provide any additional details that the connector requires to authenticate your request.
	TokenUrlCustomProperties map[string]string `json:"tokenUrlCustomProperties,omitempty" yaml:"tokenUrlCustomProperties,omitempty"`
}
