package types

type Codepipeline_PipelineTriggerGitConfigurationPush struct {
	// The field that contains the details for the Git tags trigger configuration. A `tags` block is documented below.
	Tags Codepipeline_PipelineTriggerGitConfigurationPushTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The field that specifies to filter on branches for the push trigger configuration. A `branches` block is documented below.
	Branches Codepipeline_PipelineTriggerGitConfigurationPushBranches `json:"branches,omitempty" yaml:"branches,omitempty"`

	// The field that specifies to filter on file paths for the push trigger configuration. A `file_paths` block is documented below.
	FilePaths Codepipeline_PipelineTriggerGitConfigurationPushFilePaths `json:"filePaths,omitempty" yaml:"filePaths,omitempty"`
}
