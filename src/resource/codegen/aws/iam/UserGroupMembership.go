package iam

type UserGroupMembership struct {
	// A list of IAM Groups to add the user to
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	// The name of the IAM User to add to groups
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
