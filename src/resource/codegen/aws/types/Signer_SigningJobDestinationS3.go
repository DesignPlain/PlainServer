package types

type Signer_SigningJobDestinationS3 struct {
	//
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// An Amazon S3 object key prefix that you can use to limit signed objects keys to begin with the specified prefix.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
