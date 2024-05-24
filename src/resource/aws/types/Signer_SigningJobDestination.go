package types

type Signer_SigningJobDestination struct {
	// A configuration block describing the S3 Destination object: See S3 Destination below for details.
	S3 Signer_SigningJobDestinationS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
