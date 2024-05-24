package types

type Codepipeline_PipelineTrigger struct {
	// Provides the filter criteria and the source stage for the repository event that starts the pipeline. For more information, refer to the [AWS documentation](https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-filter.html). A `git_configuration` block is documented below.
	GitConfiguration Codepipeline_PipelineTriggerGitConfiguration `json:"gitConfiguration,omitempty" yaml:"gitConfiguration,omitempty"`

	// The source provider for the event. Possible value is `CodeStarSourceConnection`.
	ProviderType string `json:"providerType,omitempty" yaml:"providerType,omitempty"`
}
