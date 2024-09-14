package types

type Codepipeline_PipelineTriggerGitConfigurationPullRequestFilePaths struct {
	// A list of patterns of Git repository file paths that, when a commit is pushed, are to be excluded from starting the pipeline.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	// A list of patterns of Git repository file paths that, when a commit is pushed, are to be included as criteria that starts the pipeline.
	Includes []string `json:"includes,omitempty" yaml:"includes,omitempty"`
}
