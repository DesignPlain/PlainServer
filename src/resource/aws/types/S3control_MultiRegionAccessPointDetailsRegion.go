package types

type S3control_MultiRegionAccessPointDetailsRegion struct {
	//
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	//
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	//
	BucketAccountId string `json:"bucketAccountId,omitempty" yaml:"bucketAccountId,omitempty"`
}
