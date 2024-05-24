package types

type Ecs_TaskDefinitionVolumeEfsVolumeConfiguration struct {
	// Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server. Transit encryption must be enabled if Amazon EFS IAM authorization is used. Valid values: `ENABLED`, `DISABLED`. If this parameter is omitted, the default value of `DISABLED` is used.
	TransitEncryption string `json:"transitEncryption,omitempty" yaml:"transitEncryption,omitempty"`

	// Port to use for transit encryption. If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses.
	TransitEncryptionPort int `json:"transitEncryptionPort,omitempty" yaml:"transitEncryptionPort,omitempty"`

	// Configuration block for authorization for the Amazon EFS file system. Detailed below.
	AuthorizationConfig Ecs_TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig `json:"authorizationConfig,omitempty" yaml:"authorizationConfig,omitempty"`

	// ID of the EFS File System.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// Directory within the Amazon EFS file system to mount as the root directory inside the host. If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying / will have the same effect as omitting this parameter. This argument is ignored when using `authorization_config`.
	RootDirectory string `json:"rootDirectory,omitempty" yaml:"rootDirectory,omitempty"`
}
