package types

type Cognito_getUserGroupsGroup struct {
	// Name of the user group.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// Precedence of the user group.
	Precedence int `json:"precedence,omitempty" yaml:"precedence,omitempty"`

	// ARN of the IAM role to be associated with the user group.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Description of the user group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
