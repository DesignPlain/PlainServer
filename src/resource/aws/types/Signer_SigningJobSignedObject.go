package types

type Signer_SigningJobSignedObject struct {
	// A configuration block describing the S3 Destination object: See S3 Destination below for details.
	S3s []Signer_SigningJobSignedObjectS3 `json:"s3s,omitempty" yaml:"s3s,omitempty"`
}
