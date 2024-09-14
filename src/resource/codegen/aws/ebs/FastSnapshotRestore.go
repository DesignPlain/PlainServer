package ebs

import types "libds/aws/types"

type FastSnapshotRestore struct {
	// Availability zone in which to enable fast snapshot restores.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// ID of the snapshot.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	//
	Timeouts types.Ebs_FastSnapshotRestoreTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
