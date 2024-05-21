package types

type Compute_NodeTemplateNodeTypeFlexibility struct {
	// Number of virtual CPUs to use.
	Cpus string `json:"cpus,omitempty" yaml:"cpus,omitempty"`

	/*
	   (Output)
	   Use local SSD
	*/
	LocalSsd string `json:"localSsd,omitempty" yaml:"localSsd,omitempty"`

	// Physical memory available to the node, defined in MB.
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}
