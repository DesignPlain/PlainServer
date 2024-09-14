package types

type Mskconnect_CustomPluginLocationS3 struct {
	// The file key for an object in an S3 bucket.
	FileKey string `json:"fileKey,omitempty" yaml:"fileKey,omitempty"`

	// The version of an object in an S3 bucket.
	ObjectVersion string `json:"objectVersion,omitempty" yaml:"objectVersion,omitempty"`

	// The Amazon Resource Name (ARN) of an S3 bucket.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`
}
