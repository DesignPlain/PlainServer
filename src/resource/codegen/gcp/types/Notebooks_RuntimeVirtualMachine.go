package types

type Notebooks_RuntimeVirtualMachine struct {
	/*
	   (Output)
	   The unique identifier of the Managed Compute Engine instance.
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   (Output)
	   The user-friendly name of the Managed Compute Engine instance.
	*/
	InstanceName string `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`

	/*
	   Virtual Machine configuration settings.
	   Structure is documented below.
	*/
	VirtualMachineConfig Notebooks_RuntimeVirtualMachineVirtualMachineConfig `json:"virtualMachineConfig,omitempty" yaml:"virtualMachineConfig,omitempty"`
}
