package securityhub

type Member struct {
	// The ID of the member AWS account.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The email of the member AWS account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
	Invite bool `json:"invite,omitempty" yaml:"invite,omitempty"`
}
