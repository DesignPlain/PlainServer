package ecr

type RepositoryPolicy struct {
	// The policy document. This is a JSON formatted string.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Name of the repository to apply the policy.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
