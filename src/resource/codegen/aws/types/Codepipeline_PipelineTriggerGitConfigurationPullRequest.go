package types

type Codepipeline_PipelineTriggerGitConfigurationPullRequest struct {
	// The field that specifies to filter on branches for the pull request trigger configuration. A `branches` block is documented below.
	Branches Codepipeline_PipelineTriggerGitConfigurationPullRequestBranches `json:"branches,omitempty" yaml:"branches,omitempty"`

	// A list that specifies which pull request events to filter on (opened, updated, closed) for the trigger configuration. Possible values are `OPEN`, `UPDATED ` and `CLOSED`.
	Events []string `json:"events,omitempty" yaml:"events,omitempty"`

	// The field that specifies to filter on file paths for the pull request trigger configuration. A `file_paths` block is documented below.
	FilePaths Codepipeline_PipelineTriggerGitConfigurationPullRequestFilePaths `json:"filePaths,omitempty" yaml:"filePaths,omitempty"`
}
