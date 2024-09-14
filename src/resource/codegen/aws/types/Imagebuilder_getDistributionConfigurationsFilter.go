package types

type Imagebuilder_getDistributionConfigurationsFilter struct {
	// Name of the filter field. Valid values can be found in the [Image Builder ListDistributionConfigurations API Reference](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_ListDistributionConfigurations.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
