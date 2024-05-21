package types

type Cloudbuild_getTriggerGitFileSource struct {
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

	// The path of the file, with the repo root as the root of the path.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   The type of the repo, since it may not be explicit from the repo field (e.g from a URL).
	   Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"]
	*/
	RepoType string `json:"repoType,omitempty" yaml:"repoType,omitempty"`

	/*
	   The fully qualified resource name of the Repo API repository. The fully qualified resource name of the Repo API repository.
	   If unspecified, the repo from which the trigger invocation originated is assumed to be the repo from which to read the specified path.
	*/
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	/*
	   The branch, tag, arbitrary ref, or SHA version of the repo to use when resolving the
	   filename (optional). This field respects the same syntax/resolution as described here: https://git-scm.com/docs/gitrevisions
	   If unspecified, the revision from which the trigger invocation originated is assumed to be the revision from which to read the specified path.
	*/
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`

	/*
	   The URI of the repo (optional). If unspecified, the repo from which the trigger
	   invocation originated is assumed to be the repo from which to read the specified path.
	*/
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
