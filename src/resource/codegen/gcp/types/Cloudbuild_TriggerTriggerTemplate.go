package types

type Cloudbuild_TriggerTriggerTemplate struct {
	/*
	   Directory, relative to the source root, in which to run the build.
	   This must be a relative path. If a step's dir is specified and
	   is an absolute path, this value is ignored for that step's
	   execution.
	*/
	Dir string `json:"dir,omitempty" yaml:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision regex.
	InvertRegex bool `json:"invertRegex,omitempty" yaml:"invertRegex,omitempty"`

	/*
	   ID of the project that owns the Cloud Source Repository. If
	   omitted, the project ID requesting the build is assumed.
	*/
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// Name of the Cloud Source Repository. If omitted, the name "default" is assumed.
	RepoName string `json:"repoName,omitempty" yaml:"repoName,omitempty"`

	/*
	   Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
	   This field is a regular expression.
	*/
	TagName string `json:"tagName,omitempty" yaml:"tagName,omitempty"`

	/*
	   Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	   This field is a regular expression.
	*/
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
	CommitSha string `json:"commitSha,omitempty" yaml:"commitSha,omitempty"`
}
