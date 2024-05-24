package types

type Athena_DatabaseAclConfiguration struct {
	/*
	   Amazon S3 canned ACL that Athena should specify when storing query results. Valid value is `BUCKET_OWNER_FULL_CONTROL`.

	   > --NOTE:-- When Athena queries are executed, result files may be created in the specified bucket. Consider using `force_destroy` on the bucket too in order to avoid any problems when destroying the bucket.
	*/
	S3AclOption string `json:"s3AclOption,omitempty" yaml:"s3AclOption,omitempty"`
}
