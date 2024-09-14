package types

type Appstream_getImageApplicationIconS3Location struct {
	// S3 bucket of the S3 object.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// S3 key of the S3 object.
	S3Key string `json:"s3Key,omitempty" yaml:"s3Key,omitempty"`
}
