package gamelift

import types "DesignSphere_Server/src/resource/aws/types"

type Build struct {
	// Name of the build
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Operating system that the game server binaries are built to run on. Valid values: `WINDOWS_2012`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `WINDOWS_2016`, `AMAZON_LINUX_2023`.
	OperatingSystem string `json:"operatingSystem,omitempty" yaml:"operatingSystem,omitempty"`

	// Information indicating where your game build files are stored. See below.
	StorageLocation types.Gamelift_BuildStorageLocation `json:"storageLocation,omitempty" yaml:"storageLocation,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Version that is associated with this build.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
