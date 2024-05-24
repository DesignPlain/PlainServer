package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation struct {
	// The file key for the object containing the application code.
	FileKey string `json:"fileKey,omitempty" yaml:"fileKey,omitempty"`

	// The version of the object containing the application code.
	ObjectVersion string `json:"objectVersion,omitempty" yaml:"objectVersion,omitempty"`

	// The ARN for the S3 bucket containing the application code.
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`
}
