package types

type Compute_RegionCommitmentResource struct {
	/*
	   Type of resource for which this commitment applies.
	   Possible values are VCPU, MEMORY, LOCAL_SSD, and ACCELERATOR.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Name of the accelerator type resource. Applicable only when the type is ACCELERATOR.
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`

	/*
	   The amount of the resource purchased (in a type-dependent unit,
	   such as bytes). For vCPUs, this can just be an integer. For memory,
	   this must be provided in MB. Memory must be a multiple of 256 MB,
	   with up to 6.5GB of memory per every vCPU.
	*/
	Amount string `json:"amount,omitempty" yaml:"amount,omitempty"`
}
