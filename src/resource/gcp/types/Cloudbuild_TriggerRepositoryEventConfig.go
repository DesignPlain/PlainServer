package types



type Cloudbuild_TriggerRepositoryEventConfig struct {
	/*
	   Contains filter properties for matching Pull Requests.
	   Structure is documented below.
	*/
	PullRequest Cloudbuild_TriggerRepositoryEventConfigPullRequest `json:"pullRequest,omitempty" yaml:"pullRequest,omitempty"`

	/*
	   Contains filter properties for matching git pushes.
	   Structure is documented below.
	*/
	Push Cloudbuild_TriggerRepositoryEventConfigPush `json:"push,omitempty" yaml:"push,omitempty"`

	// The resource name of the Repo API resource.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
