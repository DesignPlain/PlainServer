package types

type Appengine_FlexibleAppVersionResources struct {
	// Number of CPU cores needed.
	Cpu int `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// Disk size (GB) needed.
	DiskGb int `json:"diskGb,omitempty" yaml:"diskGb,omitempty"`

	// Memory (GB) needed.
	MemoryGb float64 `json:"memoryGb,omitempty" yaml:"memoryGb,omitempty"`

	/*
	   List of ports, or port pairs, to forward from the virtual machine to the application container.
	   Structure is documented below.
	*/
	Volumes []Appengine_FlexibleAppVersionResourcesVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
