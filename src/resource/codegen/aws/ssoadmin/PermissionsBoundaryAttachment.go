package ssoadmin

import types "libds/aws/types"

type PermissionsBoundaryAttachment struct {
	// The permissions boundary policy. See below.
	PermissionsBoundary types.Ssoadmin_PermissionsBoundaryAttachmentPermissionsBoundary `json:"permissionsBoundary,omitempty" yaml:"permissionsBoundary,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn string `json:"permissionSetArn,omitempty" yaml:"permissionSetArn,omitempty"`
}
