package types

type Codepipeline_PipelineStageAction struct {
	// The namespace all output variables will be accessed from.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// A list of artifact names to output. Output artifact names must be unique within a pipeline.
	OutputArtifacts []string `json:"outputArtifacts,omitempty" yaml:"outputArtifacts,omitempty"`

	// The creator of the action being called. Possible values are `AWS`, `Custom` and `ThirdParty`.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`

	// The region in which to run the action.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// A map of the action declaration's configuration. Configurations options for action types and providers can be found in the [Pipeline Structure Reference](http://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) and [Action Structure Reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) documentation. Note: The `DetectChanges` parameter (optional, default value is true) in the `configuration` section causes CodePipeline to automatically start your pipeline upon new commits. Please refer to AWS Documentation for more details: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config.
	Configuration map[string]string `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// A list of artifact names to be worked on.
	InputArtifacts []string `json:"inputArtifacts,omitempty" yaml:"inputArtifacts,omitempty"`

	// The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the [Action Structure Reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) documentation.
	Provider string `json:"provider,omitempty" yaml:"provider,omitempty"`

	// The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The order in which actions are run.
	RunOrder int `json:"runOrder,omitempty" yaml:"runOrder,omitempty"`

	//
	TimeoutInMinutes int `json:"timeoutInMinutes,omitempty" yaml:"timeoutInMinutes,omitempty"`

	// A string that identifies the action type.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are `Approval`, `Build`, `Deploy`, `Invoke`, `Source` and `Test`.
	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	// The action declaration's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
