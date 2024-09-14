package sagemaker

import types "libds/aws/types"

type FlowDefinition struct {
	// An object containing information about the tasks the human reviewers will perform. See Human Loop Config details below.
	HumanLoopConfig types.Sagemaker_FlowDefinitionHumanLoopConfig `json:"humanLoopConfig,omitempty" yaml:"humanLoopConfig,omitempty"`

	// Container for configuring the source of human task requests. Use to specify if Amazon Rekognition or Amazon Textract is used as an integration source. See Human Loop Request Source details below.
	HumanLoopRequestSource types.Sagemaker_FlowDefinitionHumanLoopRequestSource `json:"humanLoopRequestSource,omitempty" yaml:"humanLoopRequestSource,omitempty"`

	// An object containing information about where the human review results will be uploaded. See Output Config details below.
	OutputConfig types.Sagemaker_FlowDefinitionOutputConfig `json:"outputConfig,omitempty" yaml:"outputConfig,omitempty"`

	// The Amazon Resource Name (ARN) of the role needed to call other services on your behalf.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of your flow definition.
	FlowDefinitionName string `json:"flowDefinitionName,omitempty" yaml:"flowDefinitionName,omitempty"`

	// An object containing information about the events that trigger a human workflow. See Human Loop Activation Config details below.
	HumanLoopActivationConfig types.Sagemaker_FlowDefinitionHumanLoopActivationConfig `json:"humanLoopActivationConfig,omitempty" yaml:"humanLoopActivationConfig,omitempty"`
}
