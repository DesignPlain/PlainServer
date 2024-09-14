package types

type Compute_ReservationSpecificReservationInstanceProperties struct {
	/*
	   Guest accelerator type and count.
	   Structure is documented below.
	*/
	GuestAccelerators []Compute_ReservationSpecificReservationInstancePropertiesGuestAccelerator `json:"guestAccelerators,omitempty" yaml:"guestAccelerators,omitempty"`

	/*
	   The amount of local ssd to reserve with each instance. This
	   reserves disks of type `local-ssd`.
	   Structure is documented below.
	*/
	LocalSsds []Compute_ReservationSpecificReservationInstancePropertiesLocalSsd `json:"localSsds,omitempty" yaml:"localSsds,omitempty"`

	// The name of the machine type to reserve.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	/*
	   The minimum CPU platform for the reservation. For example,
	   `"Intel Skylake"`. See
	   the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	   for information on available CPU platforms.
	*/
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`
}
