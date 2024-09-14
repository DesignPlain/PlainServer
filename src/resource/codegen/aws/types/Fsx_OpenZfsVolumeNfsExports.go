package types

type Fsx_OpenZfsVolumeNfsExports struct {
	// A list of configuration objects that contain the client and options for mounting the OpenZFS file system. Maximum of 25 items. See `client_configurations` Block below for details.
	ClientConfigurations []Fsx_OpenZfsVolumeNfsExportsClientConfiguration `json:"clientConfigurations,omitempty" yaml:"clientConfigurations,omitempty"`
}
