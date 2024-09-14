package types

type Appflow_FlowSourceFlowConfig struct {
	// Name of the connector profile. This name must be unique for each connector profile in the AWS account.
	ConnectorProfileName string `json:"connectorProfileName,omitempty" yaml:"connectorProfileName,omitempty"`

	// Type of connector, such as Salesforce, Amplitude, and so on. Valid values are `Salesforce`, `Singular`, `Slack`, `Redshift`, `S3`, `Marketo`, `Googleanalytics`, `Zendesk`, `Servicenow`, `Datadog`, `Trendmicro`, `Snowflake`, `Dynatrace`, `Infornexus`, `Amplitude`, `Veeva`, `EventBridge`, `LookoutMetrics`, `Upsolver`, `Honeycode`, `CustomerProfiles`, `SAPOData`, and `CustomConnector`.
	ConnectorType string `json:"connectorType,omitempty" yaml:"connectorType,omitempty"`

	// Defines the configuration for a scheduled incremental data pull. If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull. See Incremental Pull Config for more details.
	IncrementalPullConfig Appflow_FlowSourceFlowConfigIncrementalPullConfig `json:"incrementalPullConfig,omitempty" yaml:"incrementalPullConfig,omitempty"`

	// Information that is required to query a particular source connector. See Source Connector Properties for details.
	SourceConnectorProperties Appflow_FlowSourceFlowConfigSourceConnectorProperties `json:"sourceConnectorProperties,omitempty" yaml:"sourceConnectorProperties,omitempty"`

	// API version that the destination connector uses.
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
}
