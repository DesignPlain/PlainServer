package cognito

type UserGroup struct {
	// The description of the user group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the user group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The precedence of the user group.
	Precedence int `json:"precedence,omitempty" yaml:"precedence,omitempty"`

	// The ARN of the IAM role to be associated with the user group.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The user pool ID.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`
}
