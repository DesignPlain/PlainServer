package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileProperties struct {
	// The connector-specific credentials required when using Amplitude. See Amplitude Connector Profile Credentials for more details.
	Amplitude Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesAmplitude `json:"amplitude,omitempty" yaml:"amplitude,omitempty"`

	// Connector-specific properties required when using Datadog. See Generic Connector Profile Properties for more details.
	Datadog Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDatadog `json:"datadog,omitempty" yaml:"datadog,omitempty"`

	// The connector-specific credentials required when using Amazon Honeycode. See Honeycode Connector Profile Credentials for more details.
	Honeycode Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesHoneycode `json:"honeycode,omitempty" yaml:"honeycode,omitempty"`

	// The connector-specific properties required when using Infor Nexus. See Generic Connector Profile Properties for more details.
	InforNexus Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesInforNexus `json:"inforNexus,omitempty" yaml:"inforNexus,omitempty"`

	// Connector-specific properties required when using Marketo. See Generic Connector Profile Properties for more details.
	Marketo Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesMarketo `json:"marketo,omitempty" yaml:"marketo,omitempty"`

	// The connector-specific properties required when using ServiceNow. See Generic Connector Profile Properties for more details.
	ServiceNow Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesServiceNow `json:"serviceNow,omitempty" yaml:"serviceNow,omitempty"`

	// Connector-specific properties required when using Veeva. See Generic Connector Profile Properties for more details.
	Veeva Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesVeeva `json:"veeva,omitempty" yaml:"veeva,omitempty"`

	// The connector-specific profile properties required when using the custom connector. See Custom Connector Profile Properties for more details.
	CustomConnector Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector `json:"customConnector,omitempty" yaml:"customConnector,omitempty"`

	// The connector-specific properties required when using Dynatrace. See Generic Connector Profile Properties for more details.
	Dynatrace Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesDynatrace `json:"dynatrace,omitempty" yaml:"dynatrace,omitempty"`

	// The connector-specific credentials required when using Google Analytics. See Google Analytics Connector Profile Credentials for more details.
	GoogleAnalytics Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesGoogleAnalytics `json:"googleAnalytics,omitempty" yaml:"googleAnalytics,omitempty"`

	// The connector-specific properties required when using Salesforce. See Salesforce Connector Profile Properties for more details.
	Salesforce Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSalesforce `json:"salesforce,omitempty" yaml:"salesforce,omitempty"`

	// The connector-specific properties required when using SAPOData. See SAPOData Connector Profile Properties for more details.
	SapoData Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData `json:"sapoData,omitempty" yaml:"sapoData,omitempty"`

	// Connector-specific properties required when using Slack. See Generic Connector Profile Properties for more details.
	Slack Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSlack `json:"slack,omitempty" yaml:"slack,omitempty"`

	// The connector-specific properties required when using Snowflake. See Snowflake Connector Profile Properties for more details.
	Snowflake Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake `json:"snowflake,omitempty" yaml:"snowflake,omitempty"`

	// The connector-specific credentials required when using Trend Micro. See Trend Micro Connector Profile Credentials for more details.
	Trendmicro Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesTrendmicro `json:"trendmicro,omitempty" yaml:"trendmicro,omitempty"`

	// Connector-specific properties required when using Zendesk. See Generic Connector Profile Properties for more details.
	Zendesk Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesZendesk `json:"zendesk,omitempty" yaml:"zendesk,omitempty"`

	// Connector-specific properties required when using Amazon Redshift. See Redshift Connector Profile Properties for more details.
	Redshift Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift `json:"redshift,omitempty" yaml:"redshift,omitempty"`

	// Connector-specific credentials required when using Singular. See Singular Connector Profile Credentials for more details.
	Singular Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSingular `json:"singular,omitempty" yaml:"singular,omitempty"`
}
