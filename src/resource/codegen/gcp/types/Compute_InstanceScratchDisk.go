package types

type Compute_InstanceScratchDisk struct {
	// The disk interface to use for attaching this disk; either SCSI or NVME.
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`

	/*
	   The size of the image in gigabytes. If not specified, it
	   will inherit the size of its base image.
	*/
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	/*
	   Name with which the attached disk will be accessible
	   under `/dev/disk/by-id/google--`
	*/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`
}
