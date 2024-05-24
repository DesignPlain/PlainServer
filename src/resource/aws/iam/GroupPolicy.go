package iam

type GroupPolicy struct {
	// The IAM group to attach to the policy.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	/*
	   The name of the policy. If omitted, the provider will
	   assign a random, unique name.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates a unique name beginning with the specified
	   prefix. Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The policy document. This is a JSON formatted string.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
