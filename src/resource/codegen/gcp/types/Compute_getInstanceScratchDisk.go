package types

type Compute_getInstanceScratchDisk struct {
	// The size of the image in gigabytes.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	/*
	   Name with which the attached disk is accessible
	   under `/dev/disk/by-id/`
	*/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// The disk interface used for attaching this disk. One of `SCSI` or `NVME`.
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`
}
