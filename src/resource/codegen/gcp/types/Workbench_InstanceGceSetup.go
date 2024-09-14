package types

type Workbench_InstanceGceSetup struct {
	/*
	   The network interfaces for the VM. Supports only one interface.
	   Structure is documented below.
	*/
	NetworkInterfaces []Workbench_InstanceGceSetupNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	/*
	   A set of Shielded Instance options. See [Images using supported Shielded
	   VM features](https://cloud.google.com/compute/docs/instances/modifying-shielded-vm).
	   Not all combinations are valid.
	   Structure is documented below.
	*/
	ShieldedInstanceConfig Workbench_InstanceGceSetupShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	/*
	   The definition of a boot disk.
	   Structure is documented below.
	*/
	BootDisk Workbench_InstanceGceSetupBootDisk `json:"bootDisk,omitempty" yaml:"bootDisk,omitempty"`

	/*
	   Use a container image to start the workbench instance.
	   Structure is documented below.
	*/
	ContainerImage Workbench_InstanceGceSetupContainerImage `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`

	// Optional. If true, no external IP will be assigned to this VM instance.
	DisablePublicIp bool `json:"disablePublicIp,omitempty" yaml:"disablePublicIp,omitempty"`

	// Optional. Custom metadata to apply to this instance.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   The service account that serves as an identity for the VM instance. Currently supports only one service account.
	   Structure is documented below.
	*/
	ServiceAccounts []Workbench_InstanceGceSetupServiceAccount `json:"serviceAccounts,omitempty" yaml:"serviceAccounts,omitempty"`

	/*
	   Optional. The Compute Engine tags to add to instance (see [Tagging
	   instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	*/
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Definition of a custom Compute Engine virtual machine image for starting
	   a workbench instance with the environment installed directly on the VM.
	   Structure is documented below.
	*/
	VmImage Workbench_InstanceGceSetupVmImage `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`

	/*
	   The hardware accelerators used on this instance. If you use accelerators, make sure that your configuration has
	   [enough vCPUs and memory to support the `machine_type` you have selected](https://cloud.google.com/compute/docs/gpus/#gpus-list).
	   Currently supports only one accelerator configuration.
	   Structure is documented below.
	*/
	AcceleratorConfigs []Workbench_InstanceGceSetupAcceleratorConfig `json:"acceleratorConfigs,omitempty" yaml:"acceleratorConfigs,omitempty"`

	/*
	   Data disks attached to the VM instance. Currently supports only one data disk.
	   Structure is documented below.
	*/
	DataDisks Workbench_InstanceGceSetupDataDisks `json:"dataDisks,omitempty" yaml:"dataDisks,omitempty"`

	/*
	   Optional. Flag to enable ip forwarding or not, default false/off.
	   https://cloud.google.com/vpc/docs/using-routes#canipforward
	*/
	EnableIpForwarding bool `json:"enableIpForwarding,omitempty" yaml:"enableIpForwarding,omitempty"`

	// Optional. The machine type of the VM instance. https://cloud.google.com/compute/docs/machine-resource
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`
}
