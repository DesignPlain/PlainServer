package s3

import types "libds/aws/types"

type DirectoryBucket struct {
	// Name of the bucket. The name must be in the format `[bucket_name]--[azid]--x-s3`. Use the `aws.s3.BucketV2` resource to manage general purpose buckets.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Data redundancy. Valid values: `SingleAvailabilityZone`.
	DataRedundancy string `json:"dataRedundancy,omitempty" yaml:"dataRedundancy,omitempty"`

	// Boolean that indicates all objects should be deleted from the bucket -when the bucket is destroyed- so that the bucket can be destroyed without error. These objects are -not- recoverable. This only deletes objects when the bucket is destroyed, -not- when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// Bucket location. See Location below for more details.
	Location types.S3_DirectoryBucketLocation `json:"location,omitempty" yaml:"location,omitempty"`

	// Bucket type. Valid values: `Directory`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
