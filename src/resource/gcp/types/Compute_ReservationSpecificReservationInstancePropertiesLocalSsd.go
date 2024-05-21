package types

type Compute_ReservationSpecificReservationInstancePropertiesLocalSsd struct {
	/*
	   The size of the disk in base-2 GB.

	   - - -
	*/
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	/*
	   The disk interface to use for attaching this disk.
	   Default value is `SCSI`.
	   Possible values are: `SCSI`, `NVME`.
	*/
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`
}
