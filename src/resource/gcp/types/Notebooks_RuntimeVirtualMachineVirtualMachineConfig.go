package types

type Notebooks_RuntimeVirtualMachineVirtualMachineConfig struct {
	// The Compute Engine machine type used for runtimes.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	/*
	   Shielded VM Instance configuration settings.
	   Structure is documented below.
	*/
	ShieldedInstanceConfig Notebooks_RuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	/*
	   If true, runtime will only have internal IP addresses. By default,
	   runtimes are not restricted to internal IP addresses, and will
	   have ephemeral external IP addresses assigned to each vm. This
	   `internal_ip_only` restriction can only be enabled for subnetwork
	   enabled networks, and all dependencies must be configured to be
	   accessible without external IP addresses.
	*/
	InternalIpOnly bool `json:"internalIpOnly,omitempty" yaml:"internalIpOnly,omitempty"`

	/*
	   The labels to associate with this runtime. Label --keys-- must
	   contain 1 to 63 characters, and must conform to [RFC 1035]
	   (https://www.ietf.org/rfc/rfc1035.txt). Label --values-- may be
	   empty, but, if present, must contain 1 to 63 characters, and must
	   conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
	   more than 32 labels can be associated with a cluster.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   (Output)
	   The zone where the virtual machine is located.
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   (Output)
	   The Compute Engine guest attributes. (see [Project and instance
	   guest attributes](https://cloud.google.com/compute/docs/
	   storing-retrieving-metadata#guest_attributes)).
	*/
	GuestAttributes map[string]string `json:"guestAttributes,omitempty" yaml:"guestAttributes,omitempty"`

	/*
	   The Compute Engine subnetwork to be used for machine
	   communications. Cannot be specified with network. A full URL or
	   partial URI are valid. Examples:
	   - `https://www.googleapis.com/compute/v1/projects/[project_id]/
	   regions/us-east1/subnetworks/sub0`
	   - `projects/[project_id]/regions/us-east1/subnetworks/sub0`
	*/
	Subnet string `json:"subnet,omitempty" yaml:"subnet,omitempty"`

	/*
	   Data disk option configuration settings.
	   Structure is documented below.
	*/
	DataDisk Notebooks_RuntimeVirtualMachineVirtualMachineConfigDataDisk `json:"dataDisk,omitempty" yaml:"dataDisk,omitempty"`

	/*
	   Encryption settings for virtual machine data disk.
	   Structure is documented below.
	*/
	EncryptionConfig Notebooks_RuntimeVirtualMachineVirtualMachineConfigEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	/*
	   The Compute Engine metadata entries to add to virtual machine.
	   (see [Project and instance metadata](https://cloud.google.com
	   /compute/docs/storing-retrieving-metadata#project_and_instance
	   _metadata)).
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   Reserved IP Range name is used for VPC Peering. The
	   subnetwork allocation will use the range -name- if it's assigned.
	*/
	ReservedIpRange string `json:"reservedIpRange,omitempty" yaml:"reservedIpRange,omitempty"`

	/*
	   The Compute Engine tags to add to runtime (see [Tagging instances]
	   (https://cloud.google.com/compute/docs/
	   label-or-tag-resources#tags)).
	*/
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The Compute Engine accelerator configuration for this runtime.
	   Structure is documented below.
	*/
	AcceleratorConfig Notebooks_RuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig `json:"acceleratorConfig,omitempty" yaml:"acceleratorConfig,omitempty"`

	/*
	   Use a list of container images to start the notebook instance.
	   Structure is documented below.
	*/
	ContainerImages []Notebooks_RuntimeVirtualMachineVirtualMachineConfigContainerImage `json:"containerImages,omitempty" yaml:"containerImages,omitempty"`

	/*
	   The Compute Engine network to be used for machine communications.
	   Cannot be specified with subnetwork. If neither `network` nor
	   `subnet` is specified, the "default" network of the project is
	   used, if it exists. A full URL or partial URI. Examples:
	   - `https://www.googleapis.com/compute/v1/projects/[project_id]/
	   regions/global/default`
	   - `projects/[project_id]/regions/global/default`
	   Runtimes are managed resources inside Google Infrastructure.
	   Runtimes support the following network configurations:
	   - Google Managed Network (Network & subnet are empty)
	   - Consumer Project VPC (network & subnet are required). Requires
	   configuring Private Service Access.
	   - Shared VPC (network & subnet are required). Requires
	   configuring Private Service Access.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The type of vNIC to be used on this interface. This may be gVNIC
	   or VirtioNet.
	   Possible values are: `UNSPECIFIED_NIC_TYPE`, `VIRTIO_NET`, `GVNIC`.
	*/
	NicType string `json:"nicType,omitempty" yaml:"nicType,omitempty"`
}
