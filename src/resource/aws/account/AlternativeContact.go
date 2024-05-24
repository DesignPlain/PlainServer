package account

type AlternativeContact struct {
	// Type of the alternate contact. Allowed values are: `BILLING`, `OPERATIONS`, `SECURITY`.
	AlternateContactType string `json:"alternateContactType,omitempty" yaml:"alternateContactType,omitempty"`

	// An email address for the alternate contact.
	EmailAddress string `json:"emailAddress,omitempty" yaml:"emailAddress,omitempty"`

	// Name of the alternate contact.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Phone number for the alternate contact.
	PhoneNumber string `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`

	// Title for the alternate contact.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
