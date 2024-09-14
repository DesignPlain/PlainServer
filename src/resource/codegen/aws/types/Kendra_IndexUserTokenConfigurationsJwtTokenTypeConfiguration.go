package types

type Kendra_IndexUserTokenConfigurationsJwtTokenTypeConfiguration struct {
	// The group attribute field. Minimum length of 1. Maximum length of 100.
	GroupAttributeField string `json:"groupAttributeField,omitempty" yaml:"groupAttributeField,omitempty"`

	// The issuer of the token. Minimum length of 1. Maximum length of 65.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// The location of the key. Valid values are `URL` or `SECRET_MANAGER`
	KeyLocation string `json:"keyLocation,omitempty" yaml:"keyLocation,omitempty"`

	// The Amazon Resource Name (ARN) of the secret.
	SecretsManagerArn string `json:"secretsManagerArn,omitempty" yaml:"secretsManagerArn,omitempty"`

	// The signing key URL. Valid pattern is `^(https?|ftp|file):\/\/([^\s]-)`
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The user name attribute field. Minimum length of 1. Maximum length of 100.
	UserNameAttributeField string `json:"userNameAttributeField,omitempty" yaml:"userNameAttributeField,omitempty"`

	// The regular expression that identifies the claim. Minimum length of 1. Maximum length of 100.
	ClaimRegex string `json:"claimRegex,omitempty" yaml:"claimRegex,omitempty"`
}
