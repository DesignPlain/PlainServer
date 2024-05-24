package types

type Kendra_ExperienceConfigurationContentSourceConfiguration struct {
	// The identifiers of the data sources you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	DataSourceIds []string `json:"dataSourceIds,omitempty" yaml:"dataSourceIds,omitempty"`

	// Whether to use documents you indexed directly using the `BatchPutDocument API`. Defaults to `false`.
	DirectPutContent bool `json:"directPutContent,omitempty" yaml:"directPutContent,omitempty"`

	// The identifier of the FAQs that you want to use for your Amazon Kendra experience. Maximum number of 100 items.
	FaqIds []string `json:"faqIds,omitempty" yaml:"faqIds,omitempty"`
}
