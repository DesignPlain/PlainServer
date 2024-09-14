package types

type Cognito_getUserPoolAdminCreateUserConfig struct {
	// - Whether only admins can create users.
	AllowAdminCreateUserOnly bool `json:"allowAdminCreateUserOnly,omitempty" yaml:"allowAdminCreateUserOnly,omitempty"`

	//
	InviteMessageTemplates []Cognito_getUserPoolAdminCreateUserConfigInviteMessageTemplate `json:"inviteMessageTemplates,omitempty" yaml:"inviteMessageTemplates,omitempty"`

	/*
	   - Number of days an unconfirmed user account remains valid.
	   - invite_message_template - Templates for invitation messages.
	*/
	UnusedAccountValidityDays int `json:"unusedAccountValidityDays,omitempty" yaml:"unusedAccountValidityDays,omitempty"`
}
