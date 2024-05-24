package ecrpublic

type RepositoryPolicy struct {
	// The policy document. This is a JSON formatted string.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Name of the repository to apply the policy.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`
}
