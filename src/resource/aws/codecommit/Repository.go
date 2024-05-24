package codecommit

type Repository struct {
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch string `json:"defaultBranch,omitempty" yaml:"defaultBranch,omitempty"`

	// The description of the repository. This needs to be less than 1000 characters
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ARN of the encryption key. If no key is specified, the default `aws/codecommit`` Amazon Web Services managed key is used.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
