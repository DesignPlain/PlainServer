package types

type Compute_getInstanceGuestAccelerator struct {
	// The accelerator type resource exposed to this instance. E.g. `nvidia-tesla-k80`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The number of the guest accelerator cards exposed to this instance.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`
}
