package types

type Quicksight_IamPolicyAssignmentIdentities struct {
	// Array of Quicksight group names to assign the policy to.
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	// Array of Quicksight user names to assign the policy to.
	Users []string `json:"users,omitempty" yaml:"users,omitempty"`
}
