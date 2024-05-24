package types

type Ecs_TaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration struct {
	// Configuration block for authorization for the Amazon FSx for Windows File Server file system detailed below.
	AuthorizationConfig Ecs_TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig `json:"authorizationConfig,omitempty" yaml:"authorizationConfig,omitempty"`

	// The Amazon FSx for Windows File Server file system ID to use.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.
	RootDirectory string `json:"rootDirectory,omitempty" yaml:"rootDirectory,omitempty"`
}
