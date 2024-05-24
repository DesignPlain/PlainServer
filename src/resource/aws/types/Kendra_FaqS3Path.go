package types

type Kendra_FaqS3Path struct {
	// The name of the S3 bucket that contains the file.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   The name of the file.

	   The following arguments are optional:
	*/
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
