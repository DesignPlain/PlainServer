package opsworks

type Permission struct {
	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSsh bool `json:"allowSsh,omitempty" yaml:"allowSsh,omitempty"`

	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo bool `json:"allowSudo,omitempty" yaml:"allowSudo,omitempty"`

	// The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iam_only`
	Level string `json:"level,omitempty" yaml:"level,omitempty"`

	// The stack to set the permissions for
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`

	// The user's IAM ARN to set permissions for
	UserArn string `json:"userArn,omitempty" yaml:"userArn,omitempty"`
}
