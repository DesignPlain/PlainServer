package types

type Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSetting struct {
	// The name of the data source that you're connecting to. Canvas currently supports OAuth for Snowflake and Salesforce Data Cloud. Valid values are `SalesforceGenie` and `Snowflake`.
	DataSourceName string `json:"dataSourceName,omitempty" yaml:"dataSourceName,omitempty"`

	// The ARN of an Amazon Web Services Secrets Manager secret that stores the credentials from your identity provider, such as the client ID and secret, authorization URL, and token URL.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// Describes whether OAuth for a data source is enabled or disabled in the Canvas application. Valid values are `ENABLED` and `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
