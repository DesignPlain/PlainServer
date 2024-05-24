package types

type Kendra_getExperienceConfigurationContentSourceConfiguration struct {
	// Identifiers of the data sources you want to use for your Amazon Kendra Experience.
	DataSourceIds []string `json:"dataSourceIds,omitempty" yaml:"dataSourceIds,omitempty"`

	// Whether to use documents you indexed directly using the `BatchPutDocument API`.
	DirectPutContent bool `json:"directPutContent,omitempty" yaml:"directPutContent,omitempty"`

	// Identifier of the FAQs that you want to use for your Amazon Kendra Experience.
	FaqIds []string `json:"faqIds,omitempty" yaml:"faqIds,omitempty"`
}
