package types

type Compute_getRegionDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`
}
