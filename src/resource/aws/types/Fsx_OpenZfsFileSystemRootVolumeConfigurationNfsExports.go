package types

type Fsx_OpenZfsFileSystemRootVolumeConfigurationNfsExports struct {
	// A list of configuration objects that contain the client and options for mounting the OpenZFS file system. Maximum of 25 items. See Client Configurations Below.
	ClientConfigurations []Fsx_OpenZfsFileSystemRootVolumeConfigurationNfsExportsClientConfiguration `json:"clientConfigurations,omitempty" yaml:"clientConfigurations,omitempty"`
}
