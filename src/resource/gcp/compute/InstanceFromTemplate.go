package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type InstanceFromTemplate struct {
	/*
	   The zone that the machine should be created in. If not
	   set, the provider zone is used.

	   In addition to these, all arguments from `gcp.compute.Instance` are supported
	   as a way to override the properties in the template. All exported attributes
	   from `gcp.compute.Instance` are likewise exported here.
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	// Whether the instance has virtual displays enabled.
	EnableDisplay bool `json:"enableDisplay,omitempty" yaml:"enableDisplay,omitempty"`

	// The networks attached to the instance.
	NetworkInterfaces []types.Compute_InstanceFromTemplateNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	// The boot disk for the instance.
	BootDisk types.Compute_InstanceFromTemplateBootDisk `json:"bootDisk,omitempty" yaml:"bootDisk,omitempty"`

	// A brief description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	DesiredStatus string `json:"desiredStatus,omitempty" yaml:"desiredStatus,omitempty"`

	// The scratch disks attached to the instance.
	ScratchDisks []types.Compute_InstanceFromTemplateScratchDisk `json:"scratchDisks,omitempty" yaml:"scratchDisks,omitempty"`

	// The list of tags attached to the instance.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Controls for advanced machine-related behavior features.
	AdvancedMachineFeatures types.Compute_InstanceFromTemplateAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	// Whether deletion protection is enabled on this instance.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []types.Compute_InstanceFromTemplateGuestAccelerator `json:"guestAccelerators,omitempty" yaml:"guestAccelerators,omitempty"`

	// Metadata key/value pairs made available within the instance.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// The minimum CPU platform specified for the VM instance.
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	/*
	   If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
	   stopping the instance without setting this field, the update will fail.
	*/
	AllowStoppingForUpdate bool `json:"allowStoppingForUpdate,omitempty" yaml:"allowStoppingForUpdate,omitempty"`

	// List of disks attached to the instance
	AttachedDisks []types.Compute_InstanceFromTemplateAttachedDisk `json:"attachedDisks,omitempty" yaml:"attachedDisks,omitempty"`

	/*
	   The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
	   to create.
	*/
	ConfidentialInstanceConfig types.Compute_InstanceFromTemplateConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty" yaml:"confidentialInstanceConfig,omitempty"`

	// A list of self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.
	ResourcePolicies string `json:"resourcePolicies,omitempty" yaml:"resourcePolicies,omitempty"`

	// The service account to attach to the instance.
	ServiceAccount types.Compute_InstanceFromTemplateServiceAccount `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// The shielded vm config being used by the instance.
	ShieldedInstanceConfig types.Compute_InstanceFromTemplateShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	/*
	   Name or self link of an instance
	   template to create the instance based on. It is recommended to reference
	   instance templates through their unique id (`self_link_unique` attribute).

	   - - -
	*/
	SourceInstanceTemplate string `json:"sourceInstanceTemplate,omitempty" yaml:"sourceInstanceTemplate,omitempty"`

	/*
	   A set of key/value label pairs assigned to the instance. --Note--: This field is non-authoritative, and will only manage
	   the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on
	   the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The machine type to create.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// Specifies the reservations that this instance can consume from.
	ReservationAffinity types.Compute_InstanceFromTemplateReservationAffinity `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`

	// Metadata startup scripts made available within the instance.
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" yaml:"metadataStartupScript,omitempty"`

	// Stores additional params passed with the request, but not persisted as part of resource payload.
	Params types.Compute_InstanceFromTemplateParams `json:"params,omitempty" yaml:"params,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
	   self_link nor project are provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	CanIpForward bool `json:"canIpForward,omitempty" yaml:"canIpForward,omitempty"`

	// The scheduling strategy being used by the instance.
	Scheduling types.Compute_InstanceFromTemplateScheduling `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`

	/*
	   A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
	   labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]-[a-z0-9]), concatenated with periods. The
	   entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	*/
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   A unique name for the resource, required by GCE.
	   Changing this forces a new resource to be created.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Configures network performance settings for the instance. If not specified, the instance will be created with its
	   default network performance configuration.
	*/
	NetworkPerformanceConfig types.Compute_InstanceFromTemplateNetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty" yaml:"networkPerformanceConfig,omitempty"`
}
