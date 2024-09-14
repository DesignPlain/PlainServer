package types

type Ssm_AssociationOutputLocation struct {
	// The S3 bucket name.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// The S3 bucket prefix. Results stored in the root if not configured.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`

	/*
	   The S3 bucket region.

	   Targets specify what instance IDs or tags to apply the document to and has these keys:
	*/
	S3Region string `json:"s3Region,omitempty" yaml:"s3Region,omitempty"`
}
