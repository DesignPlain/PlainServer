package iam

type Policy struct {
	// Description of the IAM policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the policy. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   Path in which to create the policy.
	   See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	*/
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The policy document. This is a JSON formatted string.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Map of resource tags for the IAM Policy. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
