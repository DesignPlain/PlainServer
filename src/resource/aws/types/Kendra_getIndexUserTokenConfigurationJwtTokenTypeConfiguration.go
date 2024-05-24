package types

type Kendra_getIndexUserTokenConfigurationJwtTokenTypeConfiguration struct {
	// ARN of the secret.
	SecretsManagerArn string `json:"secretsManagerArn,omitempty" yaml:"secretsManagerArn,omitempty"`

	// Signing key URL.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The user name attribute field.
	UserNameAttributeField string `json:"userNameAttributeField,omitempty" yaml:"userNameAttributeField,omitempty"`

	// Regular expression that identifies the claim.
	ClaimRegex string `json:"claimRegex,omitempty" yaml:"claimRegex,omitempty"`

	// The group attribute field.
	GroupAttributeField string `json:"groupAttributeField,omitempty" yaml:"groupAttributeField,omitempty"`

	// Issuer of the token.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// Location of the key. Valid values are `URL` or `SECRET_MANAGER`
	KeyLocation string `json:"keyLocation,omitempty" yaml:"keyLocation,omitempty"`
}
