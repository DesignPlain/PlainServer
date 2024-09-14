package types

type Organizations_getOrganizationalUnitChildAccountsAccount struct {
	// Parent identifier of the organizational units.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The friendly name of the account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The status of the account in the organization.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The Amazon Resource Name (ARN) of the account.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The email address associated with the AWS account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}
