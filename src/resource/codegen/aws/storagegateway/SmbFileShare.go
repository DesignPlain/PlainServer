package storagegateway

import types "libds/aws/types"

type SmbFileShare struct {
	// The region of the S3 buck used by the file share. Required when specifying a `vpc_endpoint_dns_name`.
	BucketRegion string `json:"bucketRegion,omitempty" yaml:"bucketRegion,omitempty"`

	// The name of the file share. Must be set if an S3 prefix name is set in `location_arn`.
	FileShareName string `json:"fileShareName,omitempty" yaml:"fileShareName,omitempty"`

	// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
	NotificationPolicy string `json:"notificationPolicy,omitempty" yaml:"notificationPolicy,omitempty"`

	// Access Control List permission for S3 objects. Defaults to `private`.
	ObjectAcl string `json:"objectAcl,omitempty" yaml:"objectAcl,omitempty"`

	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`

	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists []string `json:"invalidUserLists,omitempty" yaml:"invalidUserLists,omitempty"`

	// A list of users in the Active Directory that are allowed to access the file share. If you need to specify an Active directory group, add '@' before the name of the group. It will be set on Allowed group in AWS console. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists []string `json:"validUserLists,omitempty" yaml:"validUserLists,omitempty"`

	// The DNS name of the VPC endpoint for S3 private link.
	VpcEndpointDnsName string `json:"vpcEndpointDnsName,omitempty" yaml:"vpcEndpointDnsName,omitempty"`

	// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
	SmbAclEnabled bool `json:"smbAclEnabled,omitempty" yaml:"smbAclEnabled,omitempty"`

	// The files and folders on this share will only be visible to users with read access. Default value is `false`.
	AccessBasedEnumeration bool `json:"accessBasedEnumeration,omitempty" yaml:"accessBasedEnumeration,omitempty"`

	// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
	AuditDestinationArn string `json:"auditDestinationArn,omitempty" yaml:"auditDestinationArn,omitempty"`

	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication string `json:"authentication,omitempty" yaml:"authentication,omitempty"`

	/*
	   Refresh cache information. see `cache_attributes` Block for more details.

	   --Note:-- If you have previously included a `cache_attributes` block in your configuration, removing it will not reset the refresh cache value and the previous value will remain. You must explicitly set a new value to change it.
	*/
	CacheAttributes types.Storagegateway_SmbFileShareCacheAttributes `json:"cacheAttributes,omitempty" yaml:"cacheAttributes,omitempty"`

	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn string `json:"gatewayArn,omitempty" yaml:"gatewayArn,omitempty"`

	// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
	OplocksEnabled bool `json:"oplocksEnabled,omitempty" yaml:"oplocksEnabled,omitempty"`

	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays bool `json:"requesterPays,omitempty" yaml:"requesterPays,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	AdminUserLists []string `json:"adminUserLists,omitempty" yaml:"adminUserLists,omitempty"`

	// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
	CaseSensitivity string `json:"caseSensitivity,omitempty" yaml:"caseSensitivity,omitempty"`

	// The default [storage class](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-DefaultStorageClass) for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`.
	DefaultStorageClass string `json:"defaultStorageClass,omitempty" yaml:"defaultStorageClass,omitempty"`

	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled bool `json:"guessMimeTypeEnabled,omitempty" yaml:"guessMimeTypeEnabled,omitempty"`

	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted bool `json:"kmsEncrypted,omitempty" yaml:"kmsEncrypted,omitempty"`

	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The ARN of the backed storage used for storing file data.
	LocationArn string `json:"locationArn,omitempty" yaml:"locationArn,omitempty"`
}
