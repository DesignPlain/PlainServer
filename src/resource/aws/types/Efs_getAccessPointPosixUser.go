package types

type Efs_getAccessPointPosixUser struct {
	// Group ID
	Gid int `json:"gid,omitempty" yaml:"gid,omitempty"`

	// Secondary group IDs
	SecondaryGids []int `json:"secondaryGids,omitempty" yaml:"secondaryGids,omitempty"`

	// User Id
	Uid int `json:"uid,omitempty" yaml:"uid,omitempty"`
}
