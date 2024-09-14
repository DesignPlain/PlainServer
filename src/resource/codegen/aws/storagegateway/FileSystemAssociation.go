package storagegateway

import types "libds/aws/types"

type FileSystemAssociation struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
	LocationArn string `json:"locationArn,omitempty" yaml:"locationArn,omitempty"`

	// The password of the user credential.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// The Amazon Resource Name (ARN) of the storage used for the audit logs.
	AuditDestinationArn string `json:"auditDestinationArn,omitempty" yaml:"auditDestinationArn,omitempty"`

	// Refresh cache information. see Cache Attributes for more details.
	CacheAttributes types.Storagegateway_FileSystemAssociationCacheAttributes `json:"cacheAttributes,omitempty" yaml:"cacheAttributes,omitempty"`

	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `json:"gatewayArn,omitempty" yaml:"gatewayArn,omitempty"`
}
