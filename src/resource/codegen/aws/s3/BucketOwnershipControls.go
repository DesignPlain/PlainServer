package s3

import types "libds/aws/types"

type BucketOwnershipControls struct {
	// Name of the bucket that you want to associate this access point with.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Configuration block(s) with Ownership Controls rules. Detailed below.
	Rule types.S3_BucketOwnershipControlsRule `json:"rule,omitempty" yaml:"rule,omitempty"`
}
