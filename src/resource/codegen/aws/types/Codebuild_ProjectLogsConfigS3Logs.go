package types

type Codebuild_ProjectLogsConfigS3Logs struct {
	// Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket. By default, only the account that uploads the objects to the bucket has access to these objects. This property allows you to give the bucket owner access to these objects. Valid values are `NONE`, `READ_ONLY`, and `FULL`. your CodeBuild service role must have the `s3:PutBucketAcl` permission. This permission allows CodeBuild to modify the access control list for the bucket.
	BucketOwnerAccess string `json:"bucketOwnerAccess,omitempty" yaml:"bucketOwnerAccess,omitempty"`

	// Whether to disable encrypting S3 logs. Defaults to `false`.
	EncryptionDisabled bool `json:"encryptionDisabled,omitempty" yaml:"encryptionDisabled,omitempty"`

	// Name of the S3 bucket and the path prefix for S3 logs. Must be set if status is `ENABLED`, otherwise it must be empty.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Current status of logs in S3 for a build project. Valid values: `ENABLED`, `DISABLED`. Defaults to `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
