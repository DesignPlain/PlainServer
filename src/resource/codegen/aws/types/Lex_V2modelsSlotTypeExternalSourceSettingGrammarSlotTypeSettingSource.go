package types

type Lex_V2modelsSlotTypeExternalSourceSettingGrammarSlotTypeSettingSource struct {
	// KMS key required to decrypt the contents of the grammar, if any.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Name of the Amazon S3 bucket that contains the grammar source.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Path to the grammar in the Amazon S3 bucket.
	S3ObjectKey string `json:"s3ObjectKey,omitempty" yaml:"s3ObjectKey,omitempty"`
}
