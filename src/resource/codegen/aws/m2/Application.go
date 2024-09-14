package m2

import types "libds/aws/types"

type Application struct {
	// ARN of role for application to use to access AWS resources.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.M2_ApplicationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The application definition for this application. You can specify either inline JSON or an S3 bucket location.
	Definition types.M2_ApplicationDefinition `json:"definition,omitempty" yaml:"definition,omitempty"`

	// Description of the application.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Engine type must be `microfocus | bluage`.
	EngineType string `json:"engineType,omitempty" yaml:"engineType,omitempty"`

	// KMS Key to use for the Application.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	/*
	   Unique identifier of the application.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
