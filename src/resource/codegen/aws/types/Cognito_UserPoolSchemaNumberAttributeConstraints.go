package types

type Cognito_UserPoolSchemaNumberAttributeConstraints struct {
	// Maximum value of an attribute that is of the number data type.
	MaxValue string `json:"maxValue,omitempty" yaml:"maxValue,omitempty"`

	// Minimum value of an attribute that is of the number data type.
	MinValue string `json:"minValue,omitempty" yaml:"minValue,omitempty"`
}
