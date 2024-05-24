package datasync

import types "DesignSphere_Server/src/resource/aws/types"

type FsxOpenZfsFileSystem struct {
	// The Amazon Resource Name (ARN) for the FSx for OpenZfs file system.
	FsxFilesystemArn string `json:"fsxFilesystemArn,omitempty" yaml:"fsxFilesystemArn,omitempty"`

	// The type of protocol that DataSync uses to access your file system. See below.
	Protocol types.Datasync_FsxOpenZfsFileSystemProtocol `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for openzfs file system.
	SecurityGroupArns []string `json:"securityGroupArns,omitempty" yaml:"securityGroupArns,omitempty"`

	// Subdirectory to perform actions as source or destination. Must start with `/fsx`.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
