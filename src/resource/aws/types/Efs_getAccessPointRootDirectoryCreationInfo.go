package types

type Efs_getAccessPointRootDirectoryCreationInfo struct {
	// POSIX owner group ID
	OwnerGid int `json:"ownerGid,omitempty" yaml:"ownerGid,omitempty"`

	// POSIX owner user ID
	OwnerUid int `json:"ownerUid,omitempty" yaml:"ownerUid,omitempty"`

	// POSIX permissions mode
	Permissions string `json:"permissions,omitempty" yaml:"permissions,omitempty"`
}
