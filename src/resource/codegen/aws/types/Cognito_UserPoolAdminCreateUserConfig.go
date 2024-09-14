package types

type Cognito_UserPoolAdminCreateUserConfig struct {
	// Set to True if only the administrator is allowed to create user profiles. Set to False if users can sign themselves up via an app.
	AllowAdminCreateUserOnly bool `json:"allowAdminCreateUserOnly,omitempty" yaml:"allowAdminCreateUserOnly,omitempty"`

	// Invite message template structure. Detailed below.
	InviteMessageTemplate Cognito_UserPoolAdminCreateUserConfigInviteMessageTemplate `json:"inviteMessageTemplate,omitempty" yaml:"inviteMessageTemplate,omitempty"`
}
