package types

type Signer_SigningJobSource struct {
	// A configuration block describing the S3 Source object: See S3 Source below for details.
	S3 Signer_SigningJobSourceS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
