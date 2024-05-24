package types

type Sagemaker_DomainDefaultUserSettingsSharingSettings struct {
	// Whether to include the notebook cell output when sharing the notebook. The default is `Disabled`. Valid values are `Allowed` and `Disabled`.
	NotebookOutputOption string `json:"notebookOutputOption,omitempty" yaml:"notebookOutputOption,omitempty"`

	// When `notebook_output_option` is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.
	S3KmsKeyId string `json:"s3KmsKeyId,omitempty" yaml:"s3KmsKeyId,omitempty"`

	// When `notebook_output_option` is Allowed, the Amazon S3 bucket used to save the notebook cell output.
	S3OutputPath string `json:"s3OutputPath,omitempty" yaml:"s3OutputPath,omitempty"`
}
