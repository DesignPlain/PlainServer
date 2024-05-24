package types

type Identitystore_getGroupFilter struct {
	// Attribute path that is used to specify which attribute name to search. Currently, `DisplayName` is the only valid attribute path.
	AttributePath string `json:"attributePath,omitempty" yaml:"attributePath,omitempty"`

	// Value for an attribute.
	AttributeValue string `json:"attributeValue,omitempty" yaml:"attributeValue,omitempty"`
}
