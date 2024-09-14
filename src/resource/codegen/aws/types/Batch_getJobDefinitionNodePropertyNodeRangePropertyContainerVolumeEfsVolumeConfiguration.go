package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerVolumeEfsVolumeConfiguration struct {
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	TransitEncryptionPort int `json:"transitEncryptionPort,omitempty" yaml:"transitEncryptionPort,omitempty"`

	// The authorization configuration details for the Amazon EFS file system.
	AuthorizationConfigs []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerVolumeEfsVolumeConfigurationAuthorizationConfig `json:"authorizationConfigs,omitempty" yaml:"authorizationConfigs,omitempty"`

	// The Amazon EFS file system ID to use.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	RootDirectory string `json:"rootDirectory,omitempty" yaml:"rootDirectory,omitempty"`

	// Determines whether to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server
	TransitEncryption string `json:"transitEncryption,omitempty" yaml:"transitEncryption,omitempty"`
}
