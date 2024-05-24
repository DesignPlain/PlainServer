package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce struct {
	// The location of the Datadog resource.
	InstanceUrl string `json:"instanceUrl,omitempty" yaml:"instanceUrl,omitempty"`

	// Indicates whether the connector profile applies to a sandbox or production environment.
	IsSandboxEnvironment bool `json:"isSandboxEnvironment,omitempty" yaml:"isSandboxEnvironment,omitempty"`
}
