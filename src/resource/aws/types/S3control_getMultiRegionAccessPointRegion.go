package types

type S3control_getMultiRegionAccessPointRegion struct {
	// The name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The AWS account ID that owns the bucket.
	BucketAccountId string `json:"bucketAccountId,omitempty" yaml:"bucketAccountId,omitempty"`

	// The name of the region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
