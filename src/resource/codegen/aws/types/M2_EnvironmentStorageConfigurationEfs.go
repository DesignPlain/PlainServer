package types

type M2_EnvironmentStorageConfigurationEfs struct {
	// Id of the EFS filesystem to mount.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// Path to mount the filesystem on, must start with `/m2/mount/`.
	MountPoint string `json:"mountPoint,omitempty" yaml:"mountPoint,omitempty"`
}
