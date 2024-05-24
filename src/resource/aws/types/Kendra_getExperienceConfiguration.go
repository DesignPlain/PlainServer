package types

type Kendra_getExperienceConfiguration struct {
	// The identifiers of your data sources and FAQs. This is the content you want to use for your Amazon Kendra Experience. Documented below.
	ContentSourceConfigurations []Kendra_getExperienceConfigurationContentSourceConfiguration `json:"contentSourceConfigurations,omitempty" yaml:"contentSourceConfigurations,omitempty"`

	// The AWS SSO field name that contains the identifiers of your users, such as their emails. Documented below.
	UserIdentityConfigurations []Kendra_getExperienceConfigurationUserIdentityConfiguration `json:"userIdentityConfigurations,omitempty" yaml:"userIdentityConfigurations,omitempty"`
}
