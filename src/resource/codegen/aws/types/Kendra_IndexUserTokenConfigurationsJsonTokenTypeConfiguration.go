package types

type Kendra_IndexUserTokenConfigurationsJsonTokenTypeConfiguration struct {
	// The group attribute field. Minimum length of 1. Maximum length of 2048.
	GroupAttributeField string `json:"groupAttributeField,omitempty" yaml:"groupAttributeField,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 2048.
	UserNameAttributeField string `json:"userNameAttributeField,omitempty" yaml:"userNameAttributeField,omitempty"`
}
