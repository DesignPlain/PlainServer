package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type DataQualityJobDefinition struct {
	// Specifies the container that runs the monitoring job. Fields are documented below.
	DataQualityAppSpecification types.Sagemaker_DataQualityJobDefinitionDataQualityAppSpecification `json:"dataQualityAppSpecification,omitempty" yaml:"dataQualityAppSpecification,omitempty"`

	// Configures the constraints and baselines for the monitoring job. Fields are documented below.
	DataQualityBaselineConfig types.Sagemaker_DataQualityJobDefinitionDataQualityBaselineConfig `json:"dataQualityBaselineConfig,omitempty" yaml:"dataQualityBaselineConfig,omitempty"`

	// A list of inputs for the monitoring job. Fields are documented below.
	DataQualityJobInput types.Sagemaker_DataQualityJobDefinitionDataQualityJobInput `json:"dataQualityJobInput,omitempty" yaml:"dataQualityJobInput,omitempty"`

	// A time limit for how long the monitoring job is allowed to run before stopping. Fields are documented below.
	StoppingCondition types.Sagemaker_DataQualityJobDefinitionStoppingCondition `json:"stoppingCondition,omitempty" yaml:"stoppingCondition,omitempty"`

	// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The output configuration for monitoring jobs. Fields are documented below.
	DataQualityJobOutputConfig types.Sagemaker_DataQualityJobDefinitionDataQualityJobOutputConfig `json:"dataQualityJobOutputConfig,omitempty" yaml:"dataQualityJobOutputConfig,omitempty"`

	// Identifies the resources to deploy for a monitoring job. Fields are documented below.
	JobResources types.Sagemaker_DataQualityJobDefinitionJobResources `json:"jobResources,omitempty" yaml:"jobResources,omitempty"`

	// The name of the data quality job definition. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies networking configuration for the monitoring job. Fields are documented below.
	NetworkConfig types.Sagemaker_DataQualityJobDefinitionNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
