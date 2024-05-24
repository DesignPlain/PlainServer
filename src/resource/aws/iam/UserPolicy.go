package iam

type UserPolicy struct {
	// The name of the policy. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The policy document. This is a JSON formatted string.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// IAM user to which to attach this policy.
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
