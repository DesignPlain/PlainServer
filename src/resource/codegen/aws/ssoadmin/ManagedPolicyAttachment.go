package ssoadmin

type ManagedPolicyAttachment struct {
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`

	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn string `json:"managedPolicyArn,omitempty" yaml:"managedPolicyArn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn string `json:"permissionSetArn,omitempty" yaml:"permissionSetArn,omitempty"`
}
