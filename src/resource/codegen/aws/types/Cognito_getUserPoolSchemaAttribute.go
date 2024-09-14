package types

type Cognito_getUserPoolSchemaAttribute struct {
	// - Whether the attribute is for developer use only.
	DeveloperOnlyAttribute bool `json:"developerOnlyAttribute,omitempty" yaml:"developerOnlyAttribute,omitempty"`

	// - Whether the attribute can be changed after user creation.
	Mutable bool `json:"mutable,omitempty" yaml:"mutable,omitempty"`

	// - Name of the attribute.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	NumberAttributeConstraints []Cognito_getUserPoolSchemaAttributeNumberAttributeConstraint `json:"numberAttributeConstraints,omitempty" yaml:"numberAttributeConstraints,omitempty"`

	/*
	   - Whether the attribute is required during user registration.
	   - number_attribute_constraints - Constraints for numeric attributes.
	   - string_attribute_constraints - Constraints for string attributes.
	*/
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	//
	StringAttributeConstraints []Cognito_getUserPoolSchemaAttributeStringAttributeConstraint `json:"stringAttributeConstraints,omitempty" yaml:"stringAttributeConstraints,omitempty"`

	// - Data type of the attribute (e.g., string, number).
	AttributeDataType string `json:"attributeDataType,omitempty" yaml:"attributeDataType,omitempty"`
}
