package types

type Customerprofiles_DomainRuleBasedMatchingAttributeTypesSelector struct {
	// The `Email` type. You can choose from `EmailAddress`, `BusinessEmailAddress` and `PersonalEmailAddress`.
	EmailAddresses []string `json:"emailAddresses,omitempty" yaml:"emailAddresses,omitempty"`

	// The `PhoneNumber` type. You can choose from `PhoneNumber`, `HomePhoneNumber`, and `MobilePhoneNumber`.
	PhoneNumbers []string `json:"phoneNumbers,omitempty" yaml:"phoneNumbers,omitempty"`

	// The `Address` type. You can choose from `Address`, `BusinessAddress`, `MaillingAddress`, and `ShippingAddress`.
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	// Configures the `AttributeMatchingModel`, you can either choose `ONE_TO_ONE` or `MANY_TO_MANY`.
	AttributeMatchingModel string `json:"attributeMatchingModel,omitempty" yaml:"attributeMatchingModel,omitempty"`
}
