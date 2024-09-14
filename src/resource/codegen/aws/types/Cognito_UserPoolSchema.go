package types

type Cognito_UserPoolSchema struct {
	// Attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
	AttributeDataType string `json:"attributeDataType,omitempty" yaml:"attributeDataType,omitempty"`

	// Whether the attribute type is developer only.
	DeveloperOnlyAttribute bool `json:"developerOnlyAttribute,omitempty" yaml:"developerOnlyAttribute,omitempty"`

	// Whether the attribute can be changed once it has been created.
	Mutable bool `json:"mutable,omitempty" yaml:"mutable,omitempty"`

	// Name of the attribute.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for the constraints for an attribute of the number type. Detailed below.
	NumberAttributeConstraints Cognito_UserPoolSchemaNumberAttributeConstraints `json:"numberAttributeConstraints,omitempty" yaml:"numberAttributeConstraints,omitempty"`

	// Whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Constraints for an attribute of the string type. Detailed below.
	StringAttributeConstraints Cognito_UserPoolSchemaStringAttributeConstraints `json:"stringAttributeConstraints,omitempty" yaml:"stringAttributeConstraints,omitempty"`
}
