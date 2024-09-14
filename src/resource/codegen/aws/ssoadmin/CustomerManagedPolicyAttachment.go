package ssoadmin

import types "libds/aws/types"

type CustomerManagedPolicyAttachment struct {
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn string `json:"permissionSetArn,omitempty" yaml:"permissionSetArn,omitempty"`

	// Specifies the name and path of a customer managed policy. See below.
	CustomerManagedPolicyReference types.Ssoadmin_CustomerManagedPolicyAttachmentCustomerManagedPolicyReference `json:"customerManagedPolicyReference,omitempty" yaml:"customerManagedPolicyReference,omitempty"`
}
