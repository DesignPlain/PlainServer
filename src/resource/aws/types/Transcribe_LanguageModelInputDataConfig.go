package types

type Transcribe_LanguageModelInputDataConfig struct {
	// S3 URI where training data is located.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`

	/*
	   S3 URI where tuning data is located.

	   The following arguments are optional:
	*/
	TuningDataS3Uri string `json:"tuningDataS3Uri,omitempty" yaml:"tuningDataS3Uri,omitempty"`

	// IAM role with access to S3 bucket.
	DataAccessRoleArn string `json:"dataAccessRoleArn,omitempty" yaml:"dataAccessRoleArn,omitempty"`
}
