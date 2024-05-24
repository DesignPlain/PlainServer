package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type Pipeline struct {
	// The display name of the pipeline.
	PipelineDisplayName string `json:"pipelineDisplayName,omitempty" yaml:"pipelineDisplayName,omitempty"`

	// The name of the pipeline.
	PipelineName string `json:"pipelineName,omitempty" yaml:"pipelineName,omitempty"`

	// The name of the Pipeline (must be unique).
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
	ParallelismConfiguration types.Sagemaker_PipelineParallelismConfiguration `json:"parallelismConfiguration,omitempty" yaml:"parallelismConfiguration,omitempty"`

	// The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
	PipelineDefinition string `json:"pipelineDefinition,omitempty" yaml:"pipelineDefinition,omitempty"`

	// The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
	PipelineDefinitionS3Location types.Sagemaker_PipelinePipelineDefinitionS3Location `json:"pipelineDefinitionS3Location,omitempty" yaml:"pipelineDefinitionS3Location,omitempty"`

	// A description of the pipeline.
	PipelineDescription string `json:"pipelineDescription,omitempty" yaml:"pipelineDescription,omitempty"`
}
