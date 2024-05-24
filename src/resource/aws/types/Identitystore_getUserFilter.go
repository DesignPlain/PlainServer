package types

type Identitystore_getUserFilter struct {
	// Attribute path that is used to specify which attribute name to search. Currently, `UserName` is the only valid attribute path.
	AttributePath string `json:"attributePath,omitempty" yaml:"attributePath,omitempty"`

	// Value for an attribute.
	AttributeValue string `json:"attributeValue,omitempty" yaml:"attributeValue,omitempty"`
}
