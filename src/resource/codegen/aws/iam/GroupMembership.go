package iam

type GroupMembership struct {
	// The name to identify the Group Membership
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of IAM User names to associate with the Group
	Users []string `json:"users,omitempty" yaml:"users,omitempty"`

	// The IAM Group name to attach the list of `users` to
	Group string `json:"group,omitempty" yaml:"group,omitempty"`
}
