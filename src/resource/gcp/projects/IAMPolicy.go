package projects

type IAMPolicy struct {
	/*
	   The `gcp.organizations.getIAMPolicy` data source that represents
	   the IAM policy that will be applied to the project. The policy will be
	   merged with any existing policy applied to the project.

	   Changing this updates the policy.

	   Deleting this removes all policies from the project, locking out users without
	   organization-level access.
	*/
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`

	/*
	   The project id of the target project. This is not
	   inferred from the provider.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
