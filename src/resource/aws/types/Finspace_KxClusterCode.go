package types

type Finspace_KxClusterCode struct {
	// Unique name for the S3 bucket.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// Full S3 path (excluding bucket) to the .zip file that contains the code to be loaded onto the cluster when it’s started.
	S3Key string `json:"s3Key,omitempty" yaml:"s3Key,omitempty"`

	// Version of an S3 Object.
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" yaml:"s3ObjectVersion,omitempty"`
}
