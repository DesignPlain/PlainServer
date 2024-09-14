package types

type M2_EnvironmentStorageConfigurationFsx struct {
	// Id of the FSX filesystem to mount.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// Path to mount the filesystem on, must start with `/m2/mount/`.
	MountPoint string `json:"mountPoint,omitempty" yaml:"mountPoint,omitempty"`
}
