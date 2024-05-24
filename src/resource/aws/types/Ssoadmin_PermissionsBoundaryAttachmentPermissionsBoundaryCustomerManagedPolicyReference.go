package types

type Ssoadmin_PermissionsBoundaryAttachmentPermissionsBoundaryCustomerManagedPolicyReference struct {
	// Name of the customer managed IAM Policy to be attached.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The path to the IAM policy to be attached. The default is `/`. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) for more information.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
