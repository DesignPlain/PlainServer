package types

type Cloudbuild_TriggerBuildSourceRepoSource struct {
	// Name of the Cloud Source Repository.
	RepoName string `json:"repoName,omitempty" yaml:"repoName,omitempty"`

	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions map[string]string `json:"substitutions,omitempty" yaml:"substitutions,omitempty"`

	/*
	   Regex matching tags to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	   The syntax of the regular expressions accepted is the syntax accepted by RE2 and
	   described at https://github.com/google/re2/wiki/Syntax
	*/
	TagName string `json:"tagName,omitempty" yaml:"tagName,omitempty"`

	/*
	   Regex matching branches to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	   The syntax of the regular expressions accepted is the syntax accepted by RE2 and
	   described at https://github.com/google/re2/wiki/Syntax
	*/
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// Explicit commit SHA to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	CommitSha string `json:"commitSha,omitempty" yaml:"commitSha,omitempty"`

	/*
	   Directory, relative to the source root, in which to run the build.
	   This must be a relative path. If a step's dir is specified and is an absolute path,
	   this value is ignored for that step's execution.
	*/
	Dir string `json:"dir,omitempty" yaml:"dir,omitempty"`

	// Only trigger a build if the revision regex does NOT match the revision regex.
	InvertRegex bool `json:"invertRegex,omitempty" yaml:"invertRegex,omitempty"`

	/*
	   ID of the project that owns the Cloud Source Repository.
	   If omitted, the project ID requesting the build is assumed.
	*/
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
