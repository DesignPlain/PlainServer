package types

type Organizations_OrganizationalUnitAccount struct {
	// ARN of the organizational unit
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Email of the account
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Identifier of the organization unit
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The name for the organizational unit
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
