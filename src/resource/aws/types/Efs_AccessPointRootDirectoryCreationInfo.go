package types

type Efs_AccessPointRootDirectoryCreationInfo struct {
	// POSIX user ID to apply to the `root_directory`.
	OwnerUid int `json:"ownerUid,omitempty" yaml:"ownerUid,omitempty"`

	// POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
	Permissions string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// POSIX group ID to apply to the `root_directory`.
	OwnerGid int `json:"ownerGid,omitempty" yaml:"ownerGid,omitempty"`
}
