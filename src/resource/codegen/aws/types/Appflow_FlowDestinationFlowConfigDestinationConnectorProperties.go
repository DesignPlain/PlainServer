package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorProperties struct {
	//
	LookoutMetrics Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesLookoutMetrics `json:"lookoutMetrics,omitempty" yaml:"lookoutMetrics,omitempty"`

	// Properties that are required to query Marketo. See Generic Destination Properties for more details.
	Marketo Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesMarketo `json:"marketo,omitempty" yaml:"marketo,omitempty"`

	// Properties that are required to query Salesforce. See Salesforce Destination Properties for more details.
	Salesforce Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce `json:"salesforce,omitempty" yaml:"salesforce,omitempty"`

	// Properties that are required to query Snowflake. See Snowflake Destination Properties for more details.
	Snowflake Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake `json:"snowflake,omitempty" yaml:"snowflake,omitempty"`

	// Properties that are required to query Upsolver. See Upsolver Destination Properties for more details.
	Upsolver Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesUpsolver `json:"upsolver,omitempty" yaml:"upsolver,omitempty"`

	// Properties that are required to query the custom Connector. See Custom Connector Destination Properties for more details.
	CustomConnector Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector `json:"customConnector,omitempty" yaml:"customConnector,omitempty"`

	// Properties that are required to query Amazon Connect Customer Profiles. See Customer Profiles Destination Properties for more details.
	CustomerProfiles Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles `json:"customerProfiles,omitempty" yaml:"customerProfiles,omitempty"`

	// Properties that are required to query Amazon Honeycode. See Generic Destination Properties for more details.
	Honeycode Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesHoneycode `json:"honeycode,omitempty" yaml:"honeycode,omitempty"`

	// Properties that are required to query SAPOData. See SAPOData Destination Properties for more details.
	SapoData Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoData `json:"sapoData,omitempty" yaml:"sapoData,omitempty"`

	// Properties that are required to query Zendesk. See Zendesk Destination Properties for more details.
	Zendesk Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesZendesk `json:"zendesk,omitempty" yaml:"zendesk,omitempty"`

	// Properties that are required to query Amazon EventBridge. See Generic Destination Properties for more details.
	EventBridge Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge `json:"eventBridge,omitempty" yaml:"eventBridge,omitempty"`

	// Properties that are required to query Amazon Redshift. See Redshift Destination Properties for more details.
	Redshift Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesRedshift `json:"redshift,omitempty" yaml:"redshift,omitempty"`

	// Properties that are required to query Amazon S3. See S3 Destination Properties for more details.
	S3 Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
