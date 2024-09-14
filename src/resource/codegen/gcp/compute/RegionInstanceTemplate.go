package compute

import types "libds/gcp/types"

type RegionInstanceTemplate struct {
	/*
	   Disks to attach to instances created from this template.
	   This can be specified multiple times for multiple disks. Structure is
	   documented below.
	*/
	Disks []types.Compute_RegionInstanceTemplateDisk `json:"disks,omitempty" yaml:"disks,omitempty"`

	// The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates a unique name beginning with the specified
	   prefix. Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// - A list of self_links of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies string `json:"resourcePolicies,omitempty" yaml:"resourcePolicies,omitempty"`

	/*
	   Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	   --Note--: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	*/
	ShieldedInstanceConfig types.Compute_RegionInstanceTemplateShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	// Tags to attach to the instance.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configure Nested Virtualisation and Simultaneous Hyper Threading on this VM. Structure is documented below
	AdvancedMachineFeatures types.Compute_RegionInstanceTemplateAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	// Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
	ConfidentialInstanceConfig types.Compute_RegionInstanceTemplateConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty" yaml:"confidentialInstanceConfig,omitempty"`

	// A brief description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A set of key/value label pairs to assign to instances
	   created from this template.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Metadata key/value pairs to make available from
	   within instances created from this template.
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   Networks to attach to instances created from
	   this template. This can be specified multiple times for multiple networks.
	   Structure is documented below.
	*/
	NetworkInterfaces []types.Compute_RegionInstanceTemplateNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	/*
	   Configures network performance settings for the instance created from the
	   template. Structure is documented below. --Note--: `machine_type`
	   must be a [supported type](https://cloud.google.com/compute/docs/networking/configure-vm-with-high-bandwidth-configuration),
	   the `image` used must include the [`GVNIC`](https://cloud.google.com/compute/docs/networking/using-gvnic#create-instance-gvnic-image)
	   in `guest-os-features`, and `network_interface.0.nic-type` must be `GVNIC`
	   in order for this setting to take effect.
	*/
	NetworkPerformanceConfig types.Compute_RegionInstanceTemplateNetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty" yaml:"networkPerformanceConfig,omitempty"`

	/*
	   A brief description to use for instances
	   created from this template.
	*/
	InstanceDescription string `json:"instanceDescription,omitempty" yaml:"instanceDescription,omitempty"`

	/*
	   Specifies the reservations that this instance can consume from.
	   Structure is documented below.
	*/
	ReservationAffinity types.Compute_RegionInstanceTemplateReservationAffinity `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`

	/*
	   The scheduling strategy to use. More details about
	   this configuration option are detailed below.
	*/
	Scheduling types.Compute_RegionInstanceTemplateScheduling `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`

	// Service account to attach to the instance. Structure is documented below.
	ServiceAccount types.Compute_RegionInstanceTemplateServiceAccount `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// A set of key/value resource manager tag pairs to bind to the instance. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty" yaml:"resourceManagerTags,omitempty"`

	/*
	   Whether to allow sending and receiving of
	   packets with non-matching source or destination IPs. This defaults to false.
	*/
	CanIpForward bool `json:"canIpForward,omitempty" yaml:"canIpForward,omitempty"`

	/*
	   Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	   --Note--: `allow_stopping_for_update` must be set to true in order to update this field.
	*/
	EnableDisplay bool `json:"enableDisplay,omitempty" yaml:"enableDisplay,omitempty"`

	// List of the type and count of accelerator cards attached to the instance. Structure documented below.
	GuestAccelerators []types.Compute_RegionInstanceTemplateGuestAccelerator `json:"guestAccelerators,omitempty" yaml:"guestAccelerators,omitempty"`

	/*
	   The machine type to create.

	   To create a machine with a [custom type](https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types) (such as extended memory), format the value like `custom-VCPUS-MEM_IN_MB` like `custom-6-20480` for 6 vCPU and 20GB of RAM.

	   - - -
	*/
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	/*
	   An alternative to using the
	   startup-script metadata key, mostly to match the compute_instance resource.
	   This replaces the startup-script metadata key on the created instance and
	   thus the two mechanisms are not allowed to be used simultaneously.
	*/
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" yaml:"metadataStartupScript,omitempty"`

	/*
	   Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
	   `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	*/
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	/*
	   The Region in which the resource belongs.
	   If region is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
