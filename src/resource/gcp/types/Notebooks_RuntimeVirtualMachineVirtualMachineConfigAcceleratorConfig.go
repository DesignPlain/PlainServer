package types

type Notebooks_RuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig struct {
	// Count of cores of this accelerator.
	CoreCount int `json:"coreCount,omitempty" yaml:"coreCount,omitempty"`

	/*
	   Accelerator model. For valid values, see
	   `https://cloud.google.com/vertex-ai/docs/workbench/reference/
	   rest/v1/projects.locations.runtimes#AcceleratorType`
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
