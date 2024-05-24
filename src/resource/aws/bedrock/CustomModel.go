package bedrock

import types "DesignSphere_Server/src/resource/aws/types"

type CustomModel struct {
	// A map of tags to assign to the customization job and custom model. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Information about the validation dataset.
	ValidationDataConfig types.Bedrock_CustomModelValidationDataConfig `json:"validationDataConfig,omitempty" yaml:"validationDataConfig,omitempty"`

	// Configuration parameters for the private Virtual Private Cloud (VPC) that contains the resources you are using for this job.
	VpcConfig types.Bedrock_CustomModelVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Name for the custom model.
	CustomModelName string `json:"customModelName,omitempty" yaml:"customModelName,omitempty"`

	// [Parameters](https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models-hp.html) related to tuning the model.
	Hyperparameters map[string]string `json:"hyperparameters,omitempty" yaml:"hyperparameters,omitempty"`

	// The customization type. Valid values: `FINE_TUNING`, `CONTINUED_PRE_TRAINING`.
	CustomizationType string `json:"customizationType,omitempty" yaml:"customizationType,omitempty"`

	// A name for the customization job.
	JobName string `json:"jobName,omitempty" yaml:"jobName,omitempty"`

	// S3 location for the output data.
	OutputDataConfig types.Bedrock_CustomModelOutputDataConfig `json:"outputDataConfig,omitempty" yaml:"outputDataConfig,omitempty"`

	// The Amazon Resource Name (ARN) of an IAM role that Bedrock can assume to perform tasks on your behalf.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	//
	Timeouts types.Bedrock_CustomModelTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Information about the training dataset.
	TrainingDataConfig types.Bedrock_CustomModelTrainingDataConfig `json:"trainingDataConfig,omitempty" yaml:"trainingDataConfig,omitempty"`

	// The Amazon Resource Name (ARN) of the base model.
	BaseModelIdentifier string `json:"baseModelIdentifier,omitempty" yaml:"baseModelIdentifier,omitempty"`

	// The custom model is encrypted at rest using this key. Specify the key ARN.
	CustomModelKmsKeyId string `json:"customModelKmsKeyId,omitempty" yaml:"customModelKmsKeyId,omitempty"`
}
