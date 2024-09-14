package types

type Shield_ProactiveEngagementEmergencyContact struct {
	// Additional notes regarding the contact.
	ContactNotes string `json:"contactNotes,omitempty" yaml:"contactNotes,omitempty"`

	// A valid email address that will be used for this contact.
	EmailAddress string `json:"emailAddress,omitempty" yaml:"emailAddress,omitempty"`

	// A phone number, starting with `+` and up to 15 digits that will be used for this contact.
	PhoneNumber string `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`
}
