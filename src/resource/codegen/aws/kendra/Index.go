package kendra

import types "libds/aws/types"

type Index struct {
	// A block that specifies the user token configuration. Detailed below.
	UserTokenConfigurations types.Kendra_IndexUserTokenConfigurations `json:"userTokenConfigurations,omitempty" yaml:"userTokenConfigurations,omitempty"`

	// The description of the Index.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The Amazon Kendra edition to use for the index. Choose `DEVELOPER_EDITION` for indexes intended for development, testing, or proof of concept. Use `ENTERPRISE_EDITION` for your production databases. Once you set the edition for an index, it can't be changed. Defaults to `ENTERPRISE_EDITION`
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`

	// A block that specifies the identifier of the AWS KMS customer managed key (CMK) that's used to encrypt data indexed by Amazon Kendra. Amazon Kendra doesn't support asymmetric CMKs. Detailed below.
	ServerSideEncryptionConfiguration types.Kendra_IndexServerSideEncryptionConfiguration `json:"serverSideEncryptionConfiguration,omitempty" yaml:"serverSideEncryptionConfiguration,omitempty"`

	// The user context policy. Valid values are `ATTRIBUTE_FILTER` or `USER_TOKEN`. For more information, refer to [UserContextPolicy](https://docs.aws.amazon.com/kendra/latest/APIReference/API_CreateIndex.html#kendra-CreateIndex-request-UserContextPolicy). Defaults to `ATTRIBUTE_FILTER`.
	UserContextPolicy string `json:"userContextPolicy,omitempty" yaml:"userContextPolicy,omitempty"`

	// A block that enables fetching access levels of groups and users from an AWS Single Sign-On identity source. To configure this, see [UserGroupResolutionConfiguration](https://docs.aws.amazon.com/kendra/latest/dg/API_UserGroupResolutionConfiguration.html). Detailed below.
	UserGroupResolutionConfiguration types.Kendra_IndexUserGroupResolutionConfiguration `json:"userGroupResolutionConfiguration,omitempty" yaml:"userGroupResolutionConfiguration,omitempty"`

	// A block that sets the number of additional document storage and query capacity units that should be used by the index. Detailed below.
	CapacityUnits types.Kendra_IndexCapacityUnits `json:"capacityUnits,omitempty" yaml:"capacityUnits,omitempty"`

	// One or more blocks that specify the configuration settings for any metadata applied to the documents in the index. Minimum number of 0 items. Maximum number of 500 items. If specified, you must define all elements, including those that are provided by default. These index fields are documented at [Amazon Kendra Index documentation](https://docs.aws.amazon.com/kendra/latest/dg/hiw-index.html). For an example resource that defines these default index fields, refer to the default example above. For an example resource that appends additional index fields, refer to the append example above. All arguments for each block must be specified. Note that blocks cannot be removed since index fields cannot be deleted. This argument is detailed below.
	DocumentMetadataConfigurationUpdates []types.Kendra_IndexDocumentMetadataConfigurationUpdate `json:"documentMetadataConfigurationUpdates,omitempty" yaml:"documentMetadataConfigurationUpdates,omitempty"`

	// Specifies the name of the Index.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// An AWS Identity and Access Management (IAM) role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics. This is also the role you use when you call the `BatchPutDocument` API to index documents from an Amazon S3 bucket.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	/*
	   Tags to apply to the Index. If configured with a provider
	   `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
