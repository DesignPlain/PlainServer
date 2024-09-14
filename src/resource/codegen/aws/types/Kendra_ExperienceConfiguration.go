package types

type Kendra_ExperienceConfiguration struct {
	// The identifiers of your data sources and FAQs. Or, you can specify that you want to use documents indexed via the `BatchPutDocument API`. The provider will only perform drift detection of its value when present in a configuration. Detailed below.
	ContentSourceConfiguration Kendra_ExperienceConfigurationContentSourceConfiguration `json:"contentSourceConfiguration,omitempty" yaml:"contentSourceConfiguration,omitempty"`

	// The AWS SSO field name that contains the identifiers of your users, such as their emails. Detailed below.
	UserIdentityConfiguration Kendra_ExperienceConfigurationUserIdentityConfiguration `json:"userIdentityConfiguration,omitempty" yaml:"userIdentityConfiguration,omitempty"`
}
