package guardduty

type Member struct {
	// Message for invitation.
	InvitationMessage string `json:"invitationMessage,omitempty" yaml:"invitationMessage,omitempty"`

	// Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
	Invite bool `json:"invite,omitempty" yaml:"invite,omitempty"`

	// AWS account ID for member account.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The detector ID of the GuardDuty account where you want to create member accounts.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// Boolean whether an email notification is sent to the accounts. Defaults to `false`.
	DisableEmailNotification bool `json:"disableEmailNotification,omitempty" yaml:"disableEmailNotification,omitempty"`

	// Email address for member account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}
