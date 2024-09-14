package types

type Athena_WorkgroupConfigurationResultConfigurationAclConfiguration struct {
	// Amazon S3 canned ACL that Athena should specify when storing query results. Valid value is `BUCKET_OWNER_FULL_CONTROL`.
	S3AclOption string `json:"s3AclOption,omitempty" yaml:"s3AclOption,omitempty"`
}
