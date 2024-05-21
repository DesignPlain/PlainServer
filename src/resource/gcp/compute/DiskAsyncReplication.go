package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type DiskAsyncReplication struct {
	// The primary disk (source of replication).
	PrimaryDisk string `json:"primaryDisk,omitempty" yaml:"primaryDisk,omitempty"`

	/*
	   The secondary disk (target of replication). You can specify only one value. Structure is documented below.

	   The `secondary_disk` block includes:
	*/
	SecondaryDisk types.Compute_DiskAsyncReplicationSecondaryDisk `json:"secondaryDisk,omitempty" yaml:"secondaryDisk,omitempty"`
}
