package types

type Ssoadmin_PermissionsBoundaryAttachmentPermissionsBoundary struct {
	// Specifies the name and path of a customer managed policy. See below.
	CustomerManagedPolicyReference Ssoadmin_PermissionsBoundaryAttachmentPermissionsBoundaryCustomerManagedPolicyReference `json:"customerManagedPolicyReference,omitempty" yaml:"customerManagedPolicyReference,omitempty"`

	// AWS-managed IAM policy ARN to use as the permissions boundary.
	ManagedPolicyArn string `json:"managedPolicyArn,omitempty" yaml:"managedPolicyArn,omitempty"`
}
