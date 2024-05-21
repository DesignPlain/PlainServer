package types

type Compute_getReservationSpecificReservationInstancePropertyLocalSsd struct {
	// The size of the disk in base-2 GB.
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	// The disk interface to use for attaching this disk. Default value: "SCSI" Possible values: ["SCSI", "NVME"]
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`
}
