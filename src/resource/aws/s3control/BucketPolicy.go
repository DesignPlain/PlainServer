package s3control

type BucketPolicy struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// JSON string of the resource policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
