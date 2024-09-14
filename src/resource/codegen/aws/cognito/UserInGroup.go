package cognito

type UserInGroup struct {
	// The name of the group to which the user is to be added.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// The user pool ID of the user and group.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// The username of the user to be added to the group.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
