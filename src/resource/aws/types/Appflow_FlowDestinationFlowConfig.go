package types

type Appflow_FlowDestinationFlowConfig struct {
	// Type of connector, such as Salesforce, Amplitude, and so on. Valid values are `Salesforce`, `Singular`, `Slack`, `Redshift`, `S3`, `Marketo`, `Googleanalytics`, `Zendesk`, `Servicenow`, `Datadog`, `Trendmicro`, `Snowflake`, `Dynatrace`, `Infornexus`, `Amplitude`, `Veeva`, `EventBridge`, `LookoutMetrics`, `Upsolver`, `Honeycode`, `CustomerProfiles`, `SAPOData`, and `CustomConnector`.
	ConnectorType string `json:"connectorType,omitempty" yaml:"connectorType,omitempty"`

	// This stores the information that is required to query a particular connector. See Destination Connector Properties for more information.
	DestinationConnectorProperties Appflow_FlowDestinationFlowConfigDestinationConnectorProperties `json:"destinationConnectorProperties,omitempty" yaml:"destinationConnectorProperties,omitempty"`

	// API version that the destination connector uses.
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`

	// Name of the connector profile. This name must be unique for each connector profile in the AWS account.
	ConnectorProfileName string `json:"connectorProfileName,omitempty" yaml:"connectorProfileName,omitempty"`
}
