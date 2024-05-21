package types

type Compute_RegionDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`
}
