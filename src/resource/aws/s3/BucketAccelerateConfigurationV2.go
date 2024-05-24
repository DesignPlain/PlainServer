package s3

type BucketAccelerateConfigurationV2 struct {
	// Name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Transfer acceleration state of the bucket. Valid values: `Enabled`, `Suspended`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
