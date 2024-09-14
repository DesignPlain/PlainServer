package ebs

type SnapshotBlockPublicAccess struct {
	// The mode in which to enable "Block public access for snapshots" for the region. Allowed values are `block-all`, `block-new-sharing`, `unblocked`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
