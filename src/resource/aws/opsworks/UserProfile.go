package opsworks

type UserProfile struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement bool `json:"allowSelfManagement,omitempty" yaml:"allowSelfManagement,omitempty"`

	// The users public key
	SshPublicKey string `json:"sshPublicKey,omitempty" yaml:"sshPublicKey,omitempty"`

	// The ssh username, with witch this user wants to log in
	SshUsername string `json:"sshUsername,omitempty" yaml:"sshUsername,omitempty"`

	// The user's IAM ARN
	UserArn string `json:"userArn,omitempty" yaml:"userArn,omitempty"`
}
