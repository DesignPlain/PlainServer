package types

type Codepipeline_PipelineTriggerGitConfiguration struct {
	// The field where the repository event that will start the pipeline is specified as pull requests. A `pull_request` block is documented below.
	PullRequests []Codepipeline_PipelineTriggerGitConfigurationPullRequest `json:"pullRequests,omitempty" yaml:"pullRequests,omitempty"`

	// The field where the repository event that will start the pipeline, such as pushing Git tags, is specified with details. A `push` block is documented below.
	Pushes []Codepipeline_PipelineTriggerGitConfigurationPush `json:"pushes,omitempty" yaml:"pushes,omitempty"`

	// The name of the pipeline source action where the trigger configuration.
	SourceActionName string `json:"sourceActionName,omitempty" yaml:"sourceActionName,omitempty"`
}
