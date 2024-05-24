package lb

type TrustStoreRevocation struct {
	// S3 Bucket name holding the client certificate CA bundle.
	RevocationsS3Bucket string `json:"revocationsS3Bucket,omitempty" yaml:"revocationsS3Bucket,omitempty"`

	// S3 object key holding the client certificate CA bundle.
	RevocationsS3Key string `json:"revocationsS3Key,omitempty" yaml:"revocationsS3Key,omitempty"`

	// Version Id of CA bundle S3 bucket object, if versioned, defaults to latest if omitted.
	RevocationsS3ObjectVersion string `json:"revocationsS3ObjectVersion,omitempty" yaml:"revocationsS3ObjectVersion,omitempty"`

	// Trust Store ARN.
	TrustStoreArn string `json:"trustStoreArn,omitempty" yaml:"trustStoreArn,omitempty"`
}
