package types

type Cloudbuild_BitbucketServerConfigConnectedRepository struct {
	// Identifier for the project storing the repository.
	ProjectKey string `json:"projectKey,omitempty" yaml:"projectKey,omitempty"`

	// Identifier for the repository.
	RepoSlug string `json:"repoSlug,omitempty" yaml:"repoSlug,omitempty"`
}
