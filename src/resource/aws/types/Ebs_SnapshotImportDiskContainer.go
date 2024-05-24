package types

type Ebs_SnapshotImportDiskContainer struct {
	// The description of the disk image being imported.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The format of the disk image being imported. One of `VHD` or `VMDK`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// The URL to the Amazon S3-based disk image being imported. It can either be a https URL (https://..) or an Amazon S3 URL (s3://..). One of `url` or `user_bucket` must be set.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The Amazon S3 bucket for the disk image. One of `url` or `user_bucket` must be set. Detailed below.
	UserBucket Ebs_SnapshotImportDiskContainerUserBucket `json:"userBucket,omitempty" yaml:"userBucket,omitempty"`
}
