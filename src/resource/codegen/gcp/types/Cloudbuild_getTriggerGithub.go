package types

type Cloudbuild_getTriggerGithub struct {
	// filter to match changes in refs, like branches or tags. Specify only one of 'pull_request' or 'push'.
	Pushes []Cloudbuild_getTriggerGithubPush `json:"pushes,omitempty" yaml:"pushes,omitempty"`

	/*
	   The resource name of the github enterprise config that should be applied to this installation.
	   For example: "projects/{$projectId}/locations/{$locationId}/githubEnterpriseConfigs/{$configId}"
	*/
	EnterpriseConfigResourceName string `json:"enterpriseConfigResourceName,omitempty" yaml:"enterpriseConfigResourceName,omitempty"`

	/*
	   Name of the repository. For example: The name for
	   https://github.com/googlecloudplatform/cloud-builders is "cloud-builders".
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Owner of the repository. For example: The owner for
	   https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
	*/
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`

	// filter to match changes in pull requests. Specify only one of 'pull_request' or 'push'.
	PullRequests []Cloudbuild_getTriggerGithubPullRequest `json:"pullRequests,omitempty" yaml:"pullRequests,omitempty"`
}
