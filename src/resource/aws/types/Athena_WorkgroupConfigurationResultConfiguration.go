package types

type Athena_WorkgroupConfigurationResultConfiguration struct {
	// That an Amazon S3 canned ACL should be set to control ownership of stored query results. See ACL Configuration below.
	AclConfiguration Athena_WorkgroupConfigurationResultConfigurationAclConfiguration `json:"aclConfiguration,omitempty" yaml:"aclConfiguration,omitempty"`

	// Configuration block with encryption settings. See Encryption Configuration below.
	EncryptionConfiguration Athena_WorkgroupConfigurationResultConfigurationEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// AWS account ID that you expect to be the owner of the Amazon S3 bucket.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Location in Amazon S3 where your query results are stored, such as `s3://path/to/query/bucket/`. For more information, see [Queries and Query Result Files](https://docs.aws.amazon.com/athena/latest/ug/querying.html).
	OutputLocation string `json:"outputLocation,omitempty" yaml:"outputLocation,omitempty"`
}
