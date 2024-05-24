package types

type Appflow_ConnectorProfileConnectorProfileConfig struct {
	// The connector-specific credentials required by each connector. See Connector Profile Credentials for more details.
	ConnectorProfileCredentials Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentials `json:"connectorProfileCredentials,omitempty" yaml:"connectorProfileCredentials,omitempty"`

	// The connector-specific properties of the profile configuration. See Connector Profile Properties for more details.
	ConnectorProfileProperties Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileProperties `json:"connectorProfileProperties,omitempty" yaml:"connectorProfileProperties,omitempty"`
}
