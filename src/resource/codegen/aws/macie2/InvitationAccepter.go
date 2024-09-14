package macie2

type InvitationAccepter struct {
	// The AWS account ID for the account that sent the invitation.
	AdministratorAccountId string `json:"administratorAccountId,omitempty" yaml:"administratorAccountId,omitempty"`
}
