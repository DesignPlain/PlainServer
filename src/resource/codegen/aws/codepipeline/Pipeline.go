package codepipeline

import types "libds/aws/types"

type Pipeline struct {
	// Type of the pipeline. Possible values are: `V1` and `V2`. Default value is `V1`.
	PipelineType string `json:"pipelineType,omitempty" yaml:"pipelineType,omitempty"`

	// A stage block. Stages are documented below.
	Stages []types.Codepipeline_PipelineStage `json:"stages,omitempty" yaml:"stages,omitempty"`

	// A trigger block. Valid only when `pipeline_type` is `V2`. Triggers are documented below.
	Triggers []types.Codepipeline_PipelineTrigger `json:"triggers,omitempty" yaml:"triggers,omitempty"`

	// A pipeline-level variable block. Valid only when `pipeline_type` is `V2`. Variable are documented below.
	Variables []types.Codepipeline_PipelineVariable `json:"variables,omitempty" yaml:"variables,omitempty"`

	/*
	   The method that the pipeline will use to handle multiple executions. The default mode is `SUPERSEDED`. For value values, refer to the [AWS documentation](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_PipelineDeclaration.html#CodePipeline-Type-PipelineDeclaration-executionMode).

	   --Note:-- `QUEUED` or `PARALLEL` mode can only be used with V2 pipelines.
	*/
	ExecutionMode string `json:"executionMode,omitempty" yaml:"executionMode,omitempty"`

	// The name of the pipeline.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// One or more artifact_store blocks. Artifact stores are documented below.
	ArtifactStores []types.Codepipeline_PipelineArtifactStore `json:"artifactStores,omitempty" yaml:"artifactStores,omitempty"`
}
