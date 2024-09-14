package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom struct {
	// A map that holds custom authentication credentials.
	CredentialsMap map[string]string `json:"credentialsMap,omitempty" yaml:"credentialsMap,omitempty"`

	// The custom authentication type that the connector uses.
	CustomAuthenticationType string `json:"customAuthenticationType,omitempty" yaml:"customAuthenticationType,omitempty"`
}
