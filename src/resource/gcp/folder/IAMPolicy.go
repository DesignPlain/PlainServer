package folder

type IAMPolicy struct {
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	/*
	   The `gcp.organizations.getIAMPolicy` data source that represents
	   the IAM policy that will be applied to the folder. The policy will be
	   merged with any existing policy applied to the folder.

	   Changing this updates the policy.

	   Deleting this removes all policies from the folder, locking out users without
	   folder-level access.
	*/
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`
}
