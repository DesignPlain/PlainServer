package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector struct {
	// The OAuth 2.0 properties required for OAuth 2.0 authentication.
	Oauth2Properties Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOauth2Properties `json:"oauth2Properties,omitempty" yaml:"oauth2Properties,omitempty"`

	// A map of properties that are required to create a profile for the custom connector.
	ProfileProperties map[string]string `json:"profileProperties,omitempty" yaml:"profileProperties,omitempty"`
}
