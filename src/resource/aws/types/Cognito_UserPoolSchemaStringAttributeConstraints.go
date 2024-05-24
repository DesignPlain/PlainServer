package types

type Cognito_UserPoolSchemaStringAttributeConstraints struct {
	// Maximum length of an attribute value of the string type.
	MaxLength string `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`

	// Minimum length of an attribute value of the string type.
	MinLength string `json:"minLength,omitempty" yaml:"minLength,omitempty"`
}
