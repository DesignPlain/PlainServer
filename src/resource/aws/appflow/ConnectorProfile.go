package appflow

import types "DesignSphere_Server/src/resource/aws/types"

type ConnectorProfile struct {
	// Indicates the connection mode and specifies whether it is public or private. Private flows use AWS PrivateLink to route data over AWS infrastructure without exposing it to the public internet. One of: `Public`, `Private`.
	ConnectionMode string `json:"connectionMode,omitempty" yaml:"connectionMode,omitempty"`

	// The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for `CustomConnector` connector type.
	ConnectorLabel string `json:"connectorLabel,omitempty" yaml:"connectorLabel,omitempty"`

	// Defines the connector-specific configuration and credentials. See Connector Profile Config for more details.
	ConnectorProfileConfig types.Appflow_ConnectorProfileConnectorProfileConfig `json:"connectorProfileConfig,omitempty" yaml:"connectorProfileConfig,omitempty"`

	// The type of connector. One of: `Amplitude`, `CustomConnector`, `CustomerProfiles`, `Datadog`, `Dynatrace`, `EventBridge`, `Googleanalytics`, `Honeycode`, `Infornexus`, `LookoutMetrics`, `Marketo`, `Redshift`, `S3`, `Salesforce`, `SAPOData`, `Servicenow`, `Singular`, `Slack`, `Snowflake`, `Trendmicro`, `Upsolver`, `Veeva`, `Zendesk`.
	ConnectorType string `json:"connectorType,omitempty" yaml:"connectorType,omitempty"`

	// ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn string `json:"kmsArn,omitempty" yaml:"kmsArn,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
