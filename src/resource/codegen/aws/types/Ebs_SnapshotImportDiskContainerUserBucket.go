package types

type Ebs_SnapshotImportDiskContainerUserBucket struct {
	// The name of the Amazon S3 bucket where the disk image is located.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// The file name of the disk image.
	S3Key string `json:"s3Key,omitempty" yaml:"s3Key,omitempty"`
}
