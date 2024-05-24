package applicationinsights

type Application struct {
	// Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
	AutoConfigEnabled bool `json:"autoConfigEnabled,omitempty" yaml:"autoConfigEnabled,omitempty"`

	// Configures all of the resources in the resource group by applying the recommended configurations.
	AutoCreate bool `json:"autoCreate,omitempty" yaml:"autoCreate,omitempty"`

	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
	CweMonitorEnabled bool `json:"cweMonitorEnabled,omitempty" yaml:"cweMonitorEnabled,omitempty"`

	// Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED`.
	GroupingType string `json:"groupingType,omitempty" yaml:"groupingType,omitempty"`

	// When set to `true`, creates opsItems for any problems detected on an application.
	OpsCenterEnabled bool `json:"opsCenterEnabled,omitempty" yaml:"opsCenterEnabled,omitempty"`

	// SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
	OpsItemSnsTopicArn string `json:"opsItemSnsTopicArn,omitempty" yaml:"opsItemSnsTopicArn,omitempty"`

	/*
	   Name of the resource group.

	   The following arguments are optional:
	*/
	ResourceGroupName string `json:"resourceGroupName,omitempty" yaml:"resourceGroupName,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
