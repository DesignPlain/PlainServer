package compute

type OrganizationSecurityPolicyAssociation struct {
	// The resource that the security policy is attached to.
	AttachmentId string `json:"attachmentId,omitempty" yaml:"attachmentId,omitempty"`

	// The name for an association.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The security policy ID of the association.


	   - - -
	*/
	PolicyId string `json:"policyId,omitempty" yaml:"policyId,omitempty"`
}
