package comprehend

import types "DesignSphere_Server/src/resource/aws/types"

type EntityRecognizer struct {
	// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
	DataAccessRoleArn string `json:"dataAccessRoleArn,omitempty" yaml:"dataAccessRoleArn,omitempty"`

	/*
	   Configuration for the training and testing data.
	   See the `input_data_config` Configuration Block section below.
	*/
	InputDataConfig types.Comprehend_EntityRecognizerInputDataConfig `json:"inputDataConfig,omitempty" yaml:"inputDataConfig,omitempty"`

	/*
	   Name for the Entity Recognizer.
	   Has a maximum length of 63 characters.
	   Can contain upper- and lower-case letters, numbers, and hypen (`-`).

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates a unique version name beginning with the specified prefix.
	   Has a maximum length of 37 characters.
	   Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	   Conflicts with `version_name`.
	*/
	VersionNamePrefix string `json:"versionNamePrefix,omitempty" yaml:"versionNamePrefix,omitempty"`

	/*
	   Configuration parameters for VPC to contain Entity Recognizer resources.
	   See the `vpc_config` Configuration Block section below.
	*/
	VpcConfig types.Comprehend_EntityRecognizerVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	/*
	   Two-letter language code for the language.
	   One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
	*/
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// The ID or ARN of a KMS Key used to encrypt trained Entity Recognizers.
	ModelKmsKeyId string `json:"modelKmsKeyId,omitempty" yaml:"modelKmsKeyId,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Name for the version of the Entity Recognizer.
	   Each version must have a unique name within the Entity Recognizer.
	   If omitted, the provider will assign a random, unique version name.
	   If explicitly set to `""`, no version name will be set.
	   Has a maximum length of 63 characters.
	   Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	   Conflicts with `version_name_prefix`.
	*/
	VersionName string `json:"versionName,omitempty" yaml:"versionName,omitempty"`

	// ID or ARN of a KMS Key used to encrypt storage volumes during job processing.
	VolumeKmsKeyId string `json:"volumeKmsKeyId,omitempty" yaml:"volumeKmsKeyId,omitempty"`
}
