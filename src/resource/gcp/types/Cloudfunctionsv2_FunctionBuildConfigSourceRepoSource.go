package types

type Cloudfunctionsv2_FunctionBuildConfigSourceRepoSource struct {
	/*
	   ID of the project that owns the Cloud Source Repository. If omitted, the
	   project ID requesting the build is assumed.
	*/
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// Name of the Cloud Source Repository.
	RepoName string `json:"repoName,omitempty" yaml:"repoName,omitempty"`

	// Regex matching tags to build.
	TagName string `json:"tagName,omitempty" yaml:"tagName,omitempty"`

	// Regex matching branches to build.
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// Regex matching tags to build.
	CommitSha string `json:"commitSha,omitempty" yaml:"commitSha,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	Dir string `json:"dir,omitempty" yaml:"dir,omitempty"`

	/*
	   Only trigger a build if the revision regex does
	   NOT match the revision regex.
	*/
	InvertRegex bool `json:"invertRegex,omitempty" yaml:"invertRegex,omitempty"`
}
