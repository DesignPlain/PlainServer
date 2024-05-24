package types

type Kendra_getThesaurusSourceS3Path struct {
	// Name of the S3 bucket that contains the file.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Name of the file.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
