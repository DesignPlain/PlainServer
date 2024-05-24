package datasync

import types "DesignSphere_Server/src/resource/aws/types"

type EfsLocation struct {
	// Specifies the Amazon Resource Name (ARN) of the access point that DataSync uses to access the Amazon EFS file system.
	AccessPointArn string `json:"accessPointArn,omitempty" yaml:"accessPointArn,omitempty"`

	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config types.Datasync_EfsLocationEc2Config `json:"ec2Config,omitempty" yaml:"ec2Config,omitempty"`

	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn string `json:"efsFileSystemArn,omitempty" yaml:"efsFileSystemArn,omitempty"`

	// Specifies an Identity and Access Management (IAM) role that DataSync assumes when mounting the Amazon EFS file system.
	FileSystemAccessRoleArn string `json:"fileSystemAccessRoleArn,omitempty" yaml:"fileSystemAccessRoleArn,omitempty"`

	// Specifies whether you want DataSync to use TLS encryption when transferring data to or from your Amazon EFS file system. Valid values are `NONE` and `TLS1_2`.
	InTransitEncryption string `json:"inTransitEncryption,omitempty" yaml:"inTransitEncryption,omitempty"`

	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
