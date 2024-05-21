package organizations

type IAMPolicy struct {
	// The organization id of the target organization.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	/*
	   The `gcp.organizations.getIAMPolicy` data source that represents
	   the IAM policy that will be applied to the organization. The policy will be
	   merged with any existing policy applied to the organization.

	   Changing this updates the policy.

	   Deleting this removes all policies from the organization, locking out users without
	   organization-level access.
	*/
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`
}
