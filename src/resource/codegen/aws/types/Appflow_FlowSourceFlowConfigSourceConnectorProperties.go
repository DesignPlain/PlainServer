package types

type Appflow_FlowSourceFlowConfigSourceConnectorProperties struct {
	// Properties that are applied when the custom connector is being used as a source. See Custom Connector Source Properties.
	CustomConnector Appflow_FlowSourceFlowConfigSourceConnectorPropertiesCustomConnector `json:"customConnector,omitempty" yaml:"customConnector,omitempty"`

	// Information that is required for querying Infor Nexus. See Generic Source Properties for more details.
	InforNexus Appflow_FlowSourceFlowConfigSourceConnectorPropertiesInforNexus `json:"inforNexus,omitempty" yaml:"inforNexus,omitempty"`

	// Information that is required for querying Amazon S3. See S3 Source Properties for more details.
	S3 Appflow_FlowSourceFlowConfigSourceConnectorPropertiesS3 `json:"s3,omitempty" yaml:"s3,omitempty"`

	// Information that is required for querying Salesforce. See Salesforce Source Properties for more details.
	Salesforce Appflow_FlowSourceFlowConfigSourceConnectorPropertiesSalesforce `json:"salesforce,omitempty" yaml:"salesforce,omitempty"`

	// Information that is required for querying ServiceNow. See Generic Source Properties for more details.
	ServiceNow Appflow_FlowSourceFlowConfigSourceConnectorPropertiesServiceNow `json:"serviceNow,omitempty" yaml:"serviceNow,omitempty"`

	// Information that is required for querying Singular. See Generic Source Properties for more details.
	Singular Appflow_FlowSourceFlowConfigSourceConnectorPropertiesSingular `json:"singular,omitempty" yaml:"singular,omitempty"`

	// Information that is required for querying Datadog. See Generic Source Properties for more details.
	Datadog Appflow_FlowSourceFlowConfigSourceConnectorPropertiesDatadog `json:"datadog,omitempty" yaml:"datadog,omitempty"`

	// Operation to be performed on the provided Trend Micro source fields. Valid values are `PROJECTION`, `EQUAL_TO`, `ADDITION`, `MULTIPLICATION`, `DIVISION`, `SUBTRACTION`, `MASK_ALL`, `MASK_FIRST_N`, `MASK_LAST_N`, `VALIDATE_NON_NULL`, `VALIDATE_NON_ZERO`, `VALIDATE_NON_NEGATIVE`, `VALIDATE_NUMERIC`, and `NO_OP`.
	Trendmicro Appflow_FlowSourceFlowConfigSourceConnectorPropertiesTrendmicro `json:"trendmicro,omitempty" yaml:"trendmicro,omitempty"`

	// Information that is required for querying Amplitude. See Generic Source Properties for more details.
	Amplitude Appflow_FlowSourceFlowConfigSourceConnectorPropertiesAmplitude `json:"amplitude,omitempty" yaml:"amplitude,omitempty"`

	// Information that is required for querying Marketo. See Generic Source Properties for more details.
	Marketo Appflow_FlowSourceFlowConfigSourceConnectorPropertiesMarketo `json:"marketo,omitempty" yaml:"marketo,omitempty"`

	// Information that is required for querying Slack. See Generic Source Properties for more details.
	Slack Appflow_FlowSourceFlowConfigSourceConnectorPropertiesSlack `json:"slack,omitempty" yaml:"slack,omitempty"`

	// Information that is required for querying Zendesk. See Generic Source Properties for more details.
	Zendesk Appflow_FlowSourceFlowConfigSourceConnectorPropertiesZendesk `json:"zendesk,omitempty" yaml:"zendesk,omitempty"`

	// Operation to be performed on the provided Dynatrace source fields. Valid values are `PROJECTION`, `BETWEEN`, `EQUAL_TO`, `ADDITION`, `MULTIPLICATION`, `DIVISION`, `SUBTRACTION`, `MASK_ALL`, `MASK_FIRST_N`, `MASK_LAST_N`, `VALIDATE_NON_NULL`, `VALIDATE_NON_ZERO`, `VALIDATE_NON_NEGATIVE`, `VALIDATE_NUMERIC`, and `NO_OP`.
	Dynatrace Appflow_FlowSourceFlowConfigSourceConnectorPropertiesDynatrace `json:"dynatrace,omitempty" yaml:"dynatrace,omitempty"`

	// Operation to be performed on the provided Google Analytics source fields. Valid values are `PROJECTION` and `BETWEEN`.
	GoogleAnalytics Appflow_FlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics `json:"googleAnalytics,omitempty" yaml:"googleAnalytics,omitempty"`

	// Information that is required for querying SAPOData as a flow source. See SAPO Source Properties for more details.
	SapoData Appflow_FlowSourceFlowConfigSourceConnectorPropertiesSapoData `json:"sapoData,omitempty" yaml:"sapoData,omitempty"`

	// Information that is required for querying Veeva. See Veeva Source Properties for more details.
	Veeva Appflow_FlowSourceFlowConfigSourceConnectorPropertiesVeeva `json:"veeva,omitempty" yaml:"veeva,omitempty"`
}
