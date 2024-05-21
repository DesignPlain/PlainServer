package types

type Compute_InstanceGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// The accelerator type resource to expose to this instance. E.g. `nvidia-tesla-k80`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
