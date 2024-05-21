package types

type Cloudbuild_getTriggerRepositoryEventConfig struct {
	// Contains filter properties for matching Pull Requests.
	PullRequests []Cloudbuild_getTriggerRepositoryEventConfigPullRequest `json:"pullRequests,omitempty" yaml:"pullRequests,omitempty"`

	// Contains filter properties for matching git pushes.
	Pushes []Cloudbuild_getTriggerRepositoryEventConfigPush `json:"pushes,omitempty" yaml:"pushes,omitempty"`

	// The resource name of the Repo API resource.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
