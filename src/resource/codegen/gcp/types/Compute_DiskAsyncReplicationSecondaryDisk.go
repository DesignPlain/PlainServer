package types

type Compute_DiskAsyncReplicationSecondaryDisk struct {
	// The secondary disk.
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`

	/*
	   Output-only. Status of replication on the secondary disk.

	   - - -
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
