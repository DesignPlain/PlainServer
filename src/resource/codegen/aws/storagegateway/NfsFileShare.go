package storagegateway

import types "libds/aws/types"

type NfsFileShare struct {
	// The region of the S3 bucket used by the file share. Required when specifying `vpc_endpoint_dns_name`.
	BucketRegion string `json:"bucketRegion,omitempty" yaml:"bucketRegion,omitempty"`

	// Refresh cache information. see Cache Attributes for more details.
	CacheAttributes types.Storagegateway_NfsFileShareCacheAttributes `json:"cacheAttributes,omitempty" yaml:"cacheAttributes,omitempty"`

	// The name of the file share. Must be set if an S3 prefix name is set in `location_arn`.
	FileShareName string `json:"fileShareName,omitempty" yaml:"fileShareName,omitempty"`

	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn string `json:"gatewayArn,omitempty" yaml:"gatewayArn,omitempty"`

	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Nested argument with file share default values. More information below. see NFS File Share Defaults for more details.
	NfsFileShareDefaults types.Storagegateway_NfsFileShareNfsFileShareDefaults `json:"nfsFileShareDefaults,omitempty" yaml:"nfsFileShareDefaults,omitempty"`

	// The ARN of the backed storage used for storing file data.
	LocationArn string `json:"locationArn,omitempty" yaml:"locationArn,omitempty"`

	// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
	NotificationPolicy string `json:"notificationPolicy,omitempty" yaml:"notificationPolicy,omitempty"`

	// Access Control List permission for S3 objects. Defaults to `private`.
	ObjectAcl string `json:"objectAcl,omitempty" yaml:"objectAcl,omitempty"`

	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`

	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays bool `json:"requesterPays,omitempty" yaml:"requesterPays,omitempty"`

	// The DNS name of the VPC endpoint for S3 PrivateLink.
	VpcEndpointDnsName string `json:"vpcEndpointDnsName,omitempty" yaml:"vpcEndpointDnsName,omitempty"`

	// The default [storage class](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-DefaultStorageClass) for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`.
	DefaultStorageClass string `json:"defaultStorageClass,omitempty" yaml:"defaultStorageClass,omitempty"`

	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled bool `json:"guessMimeTypeEnabled,omitempty" yaml:"guessMimeTypeEnabled,omitempty"`

	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
	Squash string `json:"squash,omitempty" yaml:"squash,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Amazon Resource Name (ARN) of the storage used for audit logs.
	AuditDestinationArn string `json:"auditDestinationArn,omitempty" yaml:"auditDestinationArn,omitempty"`

	// The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
	ClientLists []string `json:"clientLists,omitempty" yaml:"clientLists,omitempty"`

	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted bool `json:"kmsEncrypted,omitempty" yaml:"kmsEncrypted,omitempty"`
}
