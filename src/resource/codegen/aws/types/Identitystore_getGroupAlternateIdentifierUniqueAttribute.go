package types

type Identitystore_getGroupAlternateIdentifierUniqueAttribute struct {
	// Attribute path that is used to specify which attribute name to search. For example: `DisplayName`. Refer to the [Group data type](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html).
	AttributePath string `json:"attributePath,omitempty" yaml:"attributePath,omitempty"`

	// Value for an attribute.
	AttributeValue string `json:"attributeValue,omitempty" yaml:"attributeValue,omitempty"`
}
