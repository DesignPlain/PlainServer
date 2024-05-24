package types

type Identitystore_getUserAlternateIdentifierUniqueAttribute struct {
	// Value for an attribute.
	AttributeValue string `json:"attributeValue,omitempty" yaml:"attributeValue,omitempty"`

	// Attribute path that is used to specify which attribute name to search. For example: `UserName`. Refer to the [User data type](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html).
	AttributePath string `json:"attributePath,omitempty" yaml:"attributePath,omitempty"`
}
