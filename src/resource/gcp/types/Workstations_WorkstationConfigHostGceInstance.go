package types

type Workstations_WorkstationConfigHostGceInstance struct {
	/*
	   An accelerator card attached to the instance.
	   Structure is documented below.
	*/
	Accelerators []Workstations_WorkstationConfigHostGceInstanceAccelerator `json:"accelerators,omitempty" yaml:"accelerators,omitempty"`

	// Size of the boot disk in GB.
	BootDiskSizeGb int `json:"bootDiskSizeGb,omitempty" yaml:"bootDiskSizeGb,omitempty"`

	/*
	   A set of Compute Engine Confidential VM instance options.
	   Structure is documented below.
	*/
	ConfidentialInstanceConfig Workstations_WorkstationConfigHostGceInstanceConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty" yaml:"confidentialInstanceConfig,omitempty"`

	// Whether instances have no public IP address.
	DisablePublicIpAddresses bool `json:"disablePublicIpAddresses,omitempty" yaml:"disablePublicIpAddresses,omitempty"`

	// Number of instances to pool for faster workstation startup.
	PoolSize int `json:"poolSize,omitempty" yaml:"poolSize,omitempty"`

	// Email address of the service account that will be used on VM instances used to support this config. This service account must have permission to pull the specified container image. If not set, VMs will run without a service account, in which case the image must be publicly accessible.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Scopes to grant to the service_account. Various scopes are automatically added based on feature usage. When specified, users of workstations under this configuration must have `iam.serviceAccounts.actAs` on the service account.
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty" yaml:"serviceAccountScopes,omitempty"`

	// Whether to disable SSH access to the VM.
	DisableSsh bool `json:"disableSsh,omitempty" yaml:"disableSsh,omitempty"`

	/*
	   Whether to enable nested virtualization on the Compute Engine VMs backing the Workstations.
	   See https://cloud.google.com/workstations/docs/reference/rest/v1beta/projects.locations.workstationClusters.workstationConfigs#GceInstance.FIELDS.enable_nested_virtualization
	*/
	EnableNestedVirtualization bool `json:"enableNestedVirtualization,omitempty" yaml:"enableNestedVirtualization,omitempty"`

	// The name of a Compute Engine machine type.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	/*
	   A set of Compute Engine Shielded instance options.
	   Structure is documented below.
	*/
	ShieldedInstanceConfig Workstations_WorkstationConfigHostGceInstanceShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	// Network tags to add to the Compute Engine machines backing the Workstations.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
