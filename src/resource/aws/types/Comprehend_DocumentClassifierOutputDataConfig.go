package types

type Comprehend_DocumentClassifierOutputDataConfig struct {
	/*
	   KMS Key used to encrypt the output documents.
	   Can be a KMS Key ID, a KMS Key ARN, a KMS Alias name, or a KMS Alias ARN.
	*/
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Full path for the output documents.
	OutputS3Uri string `json:"outputS3Uri,omitempty" yaml:"outputS3Uri,omitempty"`

	/*
	   Destination path for the output documents.
	   The full path to the output file will be returned in `output_s3_uri`.
	*/
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`
}
