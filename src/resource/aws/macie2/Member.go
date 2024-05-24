package macie2

type Member struct {
	// Send an invitation to a member
	Invite bool `json:"invite,omitempty" yaml:"invite,omitempty"`

	// Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The AWS account ID for the account.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The email address for the account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
	InvitationDisableEmailNotification bool `json:"invitationDisableEmailNotification,omitempty" yaml:"invitationDisableEmailNotification,omitempty"`

	// A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
	InvitationMessage string `json:"invitationMessage,omitempty" yaml:"invitationMessage,omitempty"`
}
