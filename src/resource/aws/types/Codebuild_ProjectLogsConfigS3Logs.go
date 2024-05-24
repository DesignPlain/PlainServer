package types

type Codebuild_ProjectLogsConfigS3Logs struct {
	// Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket. By default, only the account that uploads the objects to the bucket has access to these objects. This property allows you to give the bucket owner access to these objects. Valid values are `NONE`, `READ_ONLY`, and `FULL`. your CodeBuild service role must have the `s3:PutBucketAcl` permission. This permission allows CodeBuild to modify the access control list for the bucket.
	BucketOwnerAccess string `json:"bucketOwnerAccess,omitempty" yaml:"bucketOwnerAccess,omitempty"`

	// Whether to disable encrypting output artifacts. If `type` is set to `NO_ARTIFACTS`, this value is ignored. Defaults to `false`.
	EncryptionDisabled bool `json:"encryptionDisabled,omitempty" yaml:"encryptionDisabled,omitempty"`

	// Information about the build output artifact location. If `type` is set to `CODEPIPELINE` or `NO_ARTIFACTS`, this value is ignored. If `type` is set to `S3`, this is the name of the output bucket.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Current status of logs in S3 for a build project. Valid values: `ENABLED`, `DISABLED`. Defaults to `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
