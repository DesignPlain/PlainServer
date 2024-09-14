package types

type Compute_getDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`
}
