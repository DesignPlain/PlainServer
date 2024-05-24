package s3

type BucketRequestPaymentConfigurationV2 struct {
	// Name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Specifies who pays for the download and request fees. Valid values: `BucketOwner`, `Requester`.
	Payer string `json:"payer,omitempty" yaml:"payer,omitempty"`
}
