package ssoadmin

type AccountAssignment struct {
	// The Amazon Resource Name (ARN) of the SSO Instance.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set that the admin wants to grant the principal access to.
	PermissionSetArn string `json:"permissionSetArn,omitempty" yaml:"permissionSetArn,omitempty"`

	// An identifier for an object in SSO, such as a user or group. PrincipalIds are GUIDs (For example, `f81d4fae-7dec-11d0-a765-00a0c91e6bf6`).
	PrincipalId string `json:"principalId,omitempty" yaml:"principalId,omitempty"`

	// The entity type for which the assignment will be created. Valid values: `USER`, `GROUP`.
	PrincipalType string `json:"principalType,omitempty" yaml:"principalType,omitempty"`

	// An AWS account identifier, typically a 10-12 digit string.
	TargetId string `json:"targetId,omitempty" yaml:"targetId,omitempty"`

	// The entity type for which the assignment will be created. Valid values: `AWS_ACCOUNT`.
	TargetType string `json:"targetType,omitempty" yaml:"targetType,omitempty"`
}
