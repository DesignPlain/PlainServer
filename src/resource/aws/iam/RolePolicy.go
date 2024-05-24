package iam

type RolePolicy struct {
	/*
	   The name of the role policy. If omitted, this provider will
	   assign a random, unique name.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates a unique name beginning with the specified
	   prefix. Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The inline policy document. This is a JSON formatted string. For more information about building IAM policy documents with the provider, see the AWS IAM Policy Document Guide
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The name of the IAM role to attach to the policy.
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
