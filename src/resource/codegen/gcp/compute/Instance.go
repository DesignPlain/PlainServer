package compute

import types "libds/gcp/types"

type Instance struct {
	// Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
	ConfidentialInstanceConfig types.Compute_InstanceConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty" yaml:"confidentialInstanceConfig,omitempty"`

	/*
	   Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
	   --Note--: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
	*/
	EnableDisplay bool `json:"enableDisplay,omitempty" yaml:"enableDisplay,omitempty"`

	/*
	   List of the type and count of accelerator cards attached to the instance. Structure documented below.
	   --Note:-- GPU accelerators can only be used with `on_host_maintenance` option set to TERMINATE.
	*/
	GuestAccelerators []types.Compute_InstanceGuestAccelerator `json:"guestAccelerators,omitempty" yaml:"guestAccelerators,omitempty"`

	/*
	   A map of key/value label pairs to assign to the instance.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The machine type to create.

	   --Note:-- If you want to update this value (resize the VM) after initial creation, you must set `allow_stopping_for_update` to `true`.

	   [Custom machine types](https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types) can be formatted as `custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY_MB`, e.g. `custom-6-20480` for 6 vCPU and 20GB of RAM.

	   There is a limit of 6.5 GB per CPU unless you add [extended memory](https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#extendedmemory). You must do this explicitly by adding the suffix `-ext`, e.g. `custom-2-15360-ext` for 2 vCPU and 15 GB of memory.
	*/
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	/*
	   Additional instance parameters.
	   .
	*/
	Params types.Compute_InstanceParams `json:"params,omitempty" yaml:"params,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
	AttachedDisks []types.Compute_InstanceAttachedDisk `json:"attachedDisks,omitempty" yaml:"attachedDisks,omitempty"`

	/*
	   Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
	   --Note--: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
	   --Note--: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
	*/
	ShieldedInstanceConfig types.Compute_InstanceShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	/*
	   Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
	   `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	   --Note--: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
	*/
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	// A list of network tags to attach to the instance.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The boot disk for the instance.
	   Structure is documented below.
	*/
	BootDisk types.Compute_InstanceBootDisk `json:"bootDisk,omitempty" yaml:"bootDisk,omitempty"`

	/*
	   Desired status of the instance. Either
	   `"RUNNING"` or `"TERMINATED"`.
	*/
	DesiredStatus string `json:"desiredStatus,omitempty" yaml:"desiredStatus,omitempty"`

	/*
	   Metadata key/value pairs to make available from
	   within the instance. Ssh keys attached in the Cloud Console will be removed.
	   Add them to your config in order to keep them attached to your instance. A
	   list of default metadata values (e.g. ssh-keys) can be found [here](https://cloud.google.com/compute/docs/metadata/default-metadata-values)

	   > Depending on the OS you choose for your instance, some metadata keys have
	   special functionality.  Most linux-based images will run the content of
	   `metadata.startup-script` in a shell on every boot.  At a minimum,
	   Debian, CentOS, RHEL, SLES, Container-Optimized OS, and Ubuntu images
	   support this key.  Windows instances require other keys depending on the format
	   of the script and the time you would like it to run - see [this table](https://cloud.google.com/compute/docs/startupscript#providing_a_startup_script_for_windows_instances).
	   For the convenience of the users of `metadata.startup-script`,
	   we provide a special attribute, `metadata_startup_script`, which is documented below.
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// A brief description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Service account to attach to the instance.
	   Structure is documented below.
	   --Note--: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
	*/
	ServiceAccount types.Compute_InstanceServiceAccount `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   A unique name for the resource, required by GCE.
	   Changing this forces a new resource to be created.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// - A list of self_links of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	ResourcePolicies string `json:"resourcePolicies,omitempty" yaml:"resourcePolicies,omitempty"`

	/*
	   Scratch disks to attach to the instance. This can be
	   specified multiple times for multiple scratch disks. Structure is documented below.
	*/
	ScratchDisks []types.Compute_InstanceScratchDisk `json:"scratchDisks,omitempty" yaml:"scratchDisks,omitempty"`

	/*
	   An alternative to using the
	   startup-script metadata key, except this one forces the instance to be recreated
	   (thus re-running the script) if it is changed. This replaces the startup-script
	   metadata key on the created instance and thus the two mechanisms are not
	   allowed to be used simultaneously.  Users are free to use either mechanism - the
	   only distinction is that this separate attribute will cause a recreate on
	   modification.  On import, `metadata_startup_script` will not be set - if you
	   choose to specify it you will see a diff immediately after import causing a
	   destroy/recreate operation. If importing an instance and specifying this value
	   is desired, you will need to modify your state file.
	*/
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" yaml:"metadataStartupScript,omitempty"`

	/*
	   Whether to allow sending and receiving of
	   packets with non-matching source or destination IPs.
	   This defaults to false.
	*/
	CanIpForward bool `json:"canIpForward,omitempty" yaml:"canIpForward,omitempty"`

	/*
	   Enable deletion protection on this instance. Defaults to false.
	   --Note:-- you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	/*
	   A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
	   Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
	   The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	*/
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   Networks to attach to the instance. This can
	   be specified multiple times. Structure is documented below.

	   - - -
	*/
	NetworkInterfaces []types.Compute_InstanceNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	/*
	   (Optional, Beta
	   Configures network performance settings for the instance. Structure is
	   documented below. --Note--: `machine_type` must be a [supported type](https://cloud.google.com/compute/docs/networking/configure-vm-with-high-bandwidth-configuration),
	   the `image` used must include the [`GVNIC`](https://cloud.google.com/compute/docs/networking/using-gvnic#create-instance-gvnic-image)
	   in `guest-os-features`, and `network_interface.0.nic-type` must be `GVNIC`
	   in order for this setting to take effect.
	*/
	NetworkPerformanceConfig types.Compute_InstanceNetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty" yaml:"networkPerformanceConfig,omitempty"`

	/*
	   The scheduling strategy to use. More details about
	   this configuration option are detailed below.
	*/
	Scheduling types.Compute_InstanceScheduling `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`

	// Configure Nested Virtualisation and Simultaneous Hyper Threading  on this VM. Structure is documented below
	AdvancedMachineFeatures types.Compute_InstanceAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	// The zone that the machine should be created in. If it is not provided, the provider zone is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   Specifies the reservations that this instance can consume from.
	   Structure is documented below.
	*/
	ReservationAffinity types.Compute_InstanceReservationAffinity `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`

	/*
	   If true, allows this prvider to stop the instance to update its properties.
	   If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	*/
	AllowStoppingForUpdate bool `json:"allowStoppingForUpdate,omitempty" yaml:"allowStoppingForUpdate,omitempty"`
}
