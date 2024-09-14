package types

type Cloudbuild_TriggerSourceToBuild struct {
	// The URI of the repo.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	/*
	   The full resource name of the bitbucket server config.
	   Format: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.
	*/
	BitbucketServerConfig string `json:"bitbucketServerConfig,omitempty" yaml:"bitbucketServerConfig,omitempty"`

	/*
	   The full resource name of the github enterprise config.
	   Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.
	*/
	GithubEnterpriseConfig string `json:"githubEnterpriseConfig,omitempty" yaml:"githubEnterpriseConfig,omitempty"`

	// The branch or tag to use. Must start with "refs/" (required).
	Ref string `json:"ref,omitempty" yaml:"ref,omitempty"`

	/*
	   The type of the repo, since it may not be explicit from the repo field (e.g from a URL).
	   Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER
	   Possible values are: `UNKNOWN`, `CLOUD_SOURCE_REPOSITORIES`, `GITHUB`, `BITBUCKET_SERVER`.
	*/
	RepoType string `json:"repoType,omitempty" yaml:"repoType,omitempty"`

	/*
	   The qualified resource name of the Repo API repository.
	   Either uri or repository can be specified and is required.
	*/
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
