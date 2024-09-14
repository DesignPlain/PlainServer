package types

type Sesv2_ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfiguration struct {
	// The default value of the dimension that is published to Amazon CloudWatch if you don't provide the value of the dimension when you send an email.
	DefaultDimensionValue string `json:"defaultDimensionValue,omitempty" yaml:"defaultDimensionValue,omitempty"`

	// The name of an Amazon CloudWatch dimension associated with an email sending metric.
	DimensionName string `json:"dimensionName,omitempty" yaml:"dimensionName,omitempty"`

	// The location where the Amazon SES API v2 finds the value of a dimension to publish to Amazon CloudWatch. Valid values: `MESSAGE_TAG`, `EMAIL_HEADER`, `LINK_TAG`.
	DimensionValueSource string `json:"dimensionValueSource,omitempty" yaml:"dimensionValueSource,omitempty"`
}
