package types

type Compute_InstanceFromMachineImageScratchDisk struct {
	// Name with which the attached disk is accessible under /dev/disk/by-id/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// The disk interface used for attaching this disk. One of SCSI or NVME.
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`

	// The size of the disk in gigabytes. One of 375 or 3000.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`
}
