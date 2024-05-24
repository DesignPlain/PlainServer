package types

type Datasync_S3LocationS3Config struct {
	// ARN of the IAM Role used to connect to the S3 Bucket.
	BucketAccessRoleArn string `json:"bucketAccessRoleArn,omitempty" yaml:"bucketAccessRoleArn,omitempty"`
}
