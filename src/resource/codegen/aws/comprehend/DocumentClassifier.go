package comprehend

import types "libds/aws/types"

type DocumentClassifier struct {
	/*
	   Two-letter language code for the language.
	   One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
	*/
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	/*
	   Configuration for the output results of training.
	   See the `output_data_config` Configuration Block section below.
	*/
	OutputDataConfig types.Comprehend_DocumentClassifierOutputDataConfig `json:"outputDataConfig,omitempty" yaml:"outputDataConfig,omitempty"`

	/*
	   Name for the version of the Document Classifier.
	   Each version must have a unique name within the Document Classifier.
	   If omitted, the provider will assign a random, unique version name.
	   If explicitly set to `""`, no version name will be set.
	   Has a maximum length of 63 characters.
	   Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	   Conflicts with `version_name_prefix`.
	*/
	VersionName string `json:"versionName,omitempty" yaml:"versionName,omitempty"`

	/*
	   Configuration parameters for VPC to contain Document Classifier resources.
	   See the `vpc_config` Configuration Block section below.
	*/
	VpcConfig types.Comprehend_DocumentClassifierVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Creates a unique version name beginning with the specified prefix.
	   Has a maximum length of 37 characters.
	   Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	   Conflicts with `version_name`.
	*/
	VersionNamePrefix string `json:"versionNamePrefix,omitempty" yaml:"versionNamePrefix,omitempty"`

	/*
	   KMS Key used to encrypt storage volumes during job processing.
	   Can be a KMS Key ID or a KMS Key ARN.
	*/
	VolumeKmsKeyId string `json:"volumeKmsKeyId,omitempty" yaml:"volumeKmsKeyId,omitempty"`

	// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
	DataAccessRoleArn string `json:"dataAccessRoleArn,omitempty" yaml:"dataAccessRoleArn,omitempty"`

	/*
	   Configuration for the training and testing data.
	   See the `input_data_config` Configuration Block section below.
	*/
	InputDataConfig types.Comprehend_DocumentClassifierInputDataConfig `json:"inputDataConfig,omitempty" yaml:"inputDataConfig,omitempty"`

	/*
	   The document classification mode.
	   One of `MULTI_CLASS` or `MULTI_LABEL`.
	   `MULTI_CLASS` is also known as "Single Label" in the AWS Console.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   KMS Key used to encrypt trained Document Classifiers.
	   Can be a KMS Key ID or a KMS Key ARN.
	*/
	ModelKmsKeyId string `json:"modelKmsKeyId,omitempty" yaml:"modelKmsKeyId,omitempty"`

	/*
	   Name for the Document Classifier.
	   Has a maximum length of 63 characters.
	   Can contain upper- and lower-case letters, numbers, and hypen (`-`).

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
