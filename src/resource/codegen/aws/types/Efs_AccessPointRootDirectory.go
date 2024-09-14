package types

type Efs_AccessPointRootDirectory struct {
	// Path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide `creation_info`.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// POSIX IDs and permissions to apply to the access point's Root Directory. See Creation Info below.
	CreationInfo Efs_AccessPointRootDirectoryCreationInfo `json:"creationInfo,omitempty" yaml:"creationInfo,omitempty"`
}
