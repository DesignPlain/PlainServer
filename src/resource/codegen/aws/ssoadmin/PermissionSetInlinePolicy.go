package ssoadmin

type PermissionSetInlinePolicy struct {
	// The IAM inline policy to attach to a Permission Set.
	InlinePolicy string `json:"inlinePolicy,omitempty" yaml:"inlinePolicy,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn string `json:"permissionSetArn,omitempty" yaml:"permissionSetArn,omitempty"`
}
