package ec2

type SnapshotCreateVolumePermission struct {
	// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// A snapshot ID
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`
}
