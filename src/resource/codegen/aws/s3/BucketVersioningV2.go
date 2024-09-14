package s3

import types "libds/aws/types"

type BucketVersioningV2 struct {
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
	Mfa string `json:"mfa,omitempty" yaml:"mfa,omitempty"`

	// Configuration block for the versioning parameters. See below.
	VersioningConfiguration types.S3_BucketVersioningV2VersioningConfiguration `json:"versioningConfiguration,omitempty" yaml:"versioningConfiguration,omitempty"`

	// Name of the S3 bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
