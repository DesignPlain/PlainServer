package types

type Ec2_getInstanceTypeFpga struct {
	//
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	//
	Manufacturer string `json:"manufacturer,omitempty" yaml:"manufacturer,omitempty"`

	// Size of the instance memory, in MiB.
	MemorySize int `json:"memorySize,omitempty" yaml:"memorySize,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
