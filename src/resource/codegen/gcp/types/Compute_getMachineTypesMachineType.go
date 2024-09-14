package types

type Compute_getMachineTypesMachineType struct {
	// The configuration of bundled local SSD for the machine type. Structure is documented below.
	BundledLocalSsds []Compute_getMachineTypesMachineTypeBundledLocalSsd `json:"bundledLocalSsds,omitempty" yaml:"bundledLocalSsds,omitempty"`

	// The number of virtual CPUs that are available to the instance.
	GuestCpus int `json:"guestCpus,omitempty" yaml:"guestCpus,omitempty"`

	// The maximum persistent disks allowed.
	MaximumPersistentDisks int `json:"maximumPersistentDisks,omitempty" yaml:"maximumPersistentDisks,omitempty"`

	// The maximum total persistent disks size (GB) allowed.
	MaximumPersistentDisksSizeGb int `json:"maximumPersistentDisksSizeGb,omitempty" yaml:"maximumPersistentDisksSizeGb,omitempty"`

	// The server-defined URL for the machine type.
	SelfLink string `json:"selfLink,omitempty" yaml:"selfLink,omitempty"`

	// A list of accelerator configurations assigned to this machine type. Structure is documented below.
	Accelerators []Compute_getMachineTypesMachineTypeAccelerator `json:"accelerators,omitempty" yaml:"accelerators,omitempty"`

	// The deprecation status associated with this machine type. Structure is documented below.
	Deprecateds []Compute_getMachineTypesMachineTypeDeprecated `json:"deprecateds,omitempty" yaml:"deprecateds,omitempty"`

	// A textual description of the machine type.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether this machine type has a shared CPU.
	IsSharedCpus bool `json:"isSharedCpus,omitempty" yaml:"isSharedCpus,omitempty"`

	// The amount of physical memory available to the instance, defined in MB.
	MemoryMb int `json:"memoryMb,omitempty" yaml:"memoryMb,omitempty"`

	// The name of the machine type.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
