package types

type Sesv2_ConfigurationSetEventDestinationEventDestinationCloudWatchDestination struct {
	// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch. See dimension_configuration below.
	DimensionConfigurations []Sesv2_ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfiguration `json:"dimensionConfigurations,omitempty" yaml:"dimensionConfigurations,omitempty"`
}
