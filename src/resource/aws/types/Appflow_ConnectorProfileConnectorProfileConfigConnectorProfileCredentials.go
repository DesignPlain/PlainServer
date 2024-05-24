package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentials struct {
	// The connector-specific credentials required when using Trend Micro. See Trend Micro Connector Profile Credentials for more details.
	Trendmicro Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro `json:"trendmicro,omitempty" yaml:"trendmicro,omitempty"`

	// Connector-specific credentials required when using Zendesk. See Zendesk Connector Profile Credentials for more details.
	Zendesk Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsZendesk `json:"zendesk,omitempty" yaml:"zendesk,omitempty"`

	// The connector-specific credentials required when using Amplitude. See Amplitude Connector Profile Credentials for more details.
	Amplitude Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude `json:"amplitude,omitempty" yaml:"amplitude,omitempty"`

	// The connector-specific credentials required when using Amazon Honeycode. See Honeycode Connector Profile Credentials for more details.
	Honeycode Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsHoneycode `json:"honeycode,omitempty" yaml:"honeycode,omitempty"`

	// The connector-specific credentials required when using Salesforce. See Salesforce Connector Profile Credentials for more details.
	Salesforce Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce `json:"salesforce,omitempty" yaml:"salesforce,omitempty"`

	// The connector-specific profile credentials required when using the custom connector. See Custom Connector Profile Credentials for more details.
	CustomConnector Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector `json:"customConnector,omitempty" yaml:"customConnector,omitempty"`

	// Connector-specific credentials required when using Slack. See Slack Connector Profile Credentials for more details.
	Slack Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack `json:"slack,omitempty" yaml:"slack,omitempty"`

	// The connector-specific credentials required when using Snowflake. See Snowflake Connector Profile Credentials for more details.
	Snowflake Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake `json:"snowflake,omitempty" yaml:"snowflake,omitempty"`

	// Connector-specific credentials required when using Veeva. See Veeva Connector Profile Credentials for more details.
	Veeva Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsVeeva `json:"veeva,omitempty" yaml:"veeva,omitempty"`

	// The connector-specific credentials required when using SAPOData. See SAPOData Connector Profile Credentials for more details.
	SapoData Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData `json:"sapoData,omitempty" yaml:"sapoData,omitempty"`

	// The connector-specific credentials required when using ServiceNow. See ServiceNow Connector Profile Credentials for more details.
	ServiceNow Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow `json:"serviceNow,omitempty" yaml:"serviceNow,omitempty"`

	// Connector-specific credentials required when using Singular. See Singular Connector Profile Credentials for more details.
	Singular Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSingular `json:"singular,omitempty" yaml:"singular,omitempty"`

	// The connector-specific credentials required when using Infor Nexus. See Infor Nexus Connector Profile Credentials for more details.
	InforNexus Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus `json:"inforNexus,omitempty" yaml:"inforNexus,omitempty"`

	// Connector-specific credentials required when using Marketo. See Marketo Connector Profile Credentials for more details.
	Marketo Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo `json:"marketo,omitempty" yaml:"marketo,omitempty"`

	// Connector-specific credentials required when using Amazon Redshift. See Redshift Connector Profile Credentials for more details.
	Redshift Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift `json:"redshift,omitempty" yaml:"redshift,omitempty"`

	// Connector-specific credentials required when using Datadog. See Datadog Connector Profile Credentials for more details.
	Datadog Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog `json:"datadog,omitempty" yaml:"datadog,omitempty"`

	// The connector-specific credentials required when using Dynatrace. See Dynatrace Connector Profile Credentials for more details.
	Dynatrace Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace `json:"dynatrace,omitempty" yaml:"dynatrace,omitempty"`

	// The connector-specific credentials required when using Google Analytics. See Google Analytics Connector Profile Credentials for more details.
	GoogleAnalytics Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalytics `json:"googleAnalytics,omitempty" yaml:"googleAnalytics,omitempty"`
}
