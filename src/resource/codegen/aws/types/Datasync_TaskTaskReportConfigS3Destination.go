package types

type Datasync_TaskTaskReportConfigS3Destination struct {
	// Specifies the Amazon Resource Name (ARN) of the IAM policy that allows DataSync to upload a task report to your S3 bucket.
	BucketAccessRoleArn string `json:"bucketAccessRoleArn,omitempty" yaml:"bucketAccessRoleArn,omitempty"`

	// Specifies the ARN of the S3 bucket where DataSync uploads your report.
	S3BucketArn string `json:"s3BucketArn,omitempty" yaml:"s3BucketArn,omitempty"`

	// Specifies a bucket prefix for your report.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`
}
