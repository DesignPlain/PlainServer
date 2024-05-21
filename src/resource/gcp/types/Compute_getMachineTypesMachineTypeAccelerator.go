package types

type Compute_getMachineTypesMachineTypeAccelerator struct {
	// Number of accelerator cards exposed to the guest.
	GuestAcceleratorCount int `json:"guestAcceleratorCount,omitempty" yaml:"guestAcceleratorCount,omitempty"`

	// The accelerator type resource name, not a full URL, e.g. `nvidia-tesla-t4`.
	GuestAcceleratorType string `json:"guestAcceleratorType,omitempty" yaml:"guestAcceleratorType,omitempty"`
}
