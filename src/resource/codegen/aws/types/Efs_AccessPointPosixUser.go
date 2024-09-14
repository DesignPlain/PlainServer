package types

type Efs_AccessPointPosixUser struct {
	// POSIX user ID used for all file system operations using this access point.
	Uid int `json:"uid,omitempty" yaml:"uid,omitempty"`

	// POSIX group ID used for all file system operations using this access point.
	Gid int `json:"gid,omitempty" yaml:"gid,omitempty"`

	// Secondary POSIX group IDs used for all file system operations using this access point.
	SecondaryGids []int `json:"secondaryGids,omitempty" yaml:"secondaryGids,omitempty"`
}
