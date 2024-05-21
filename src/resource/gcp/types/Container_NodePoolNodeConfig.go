package types



type Container_NodePoolNodeConfig struct {
	// The set of Google API scopes to be made available on all of the node VMs.
	OauthScopes []string `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`

	// The GCE resource labels (a map of key/value pairs) to be applied to the node pool.
	ResourceLabels map[string]string `json:"resourceLabels,omitempty" yaml:"resourceLabels,omitempty"`

	// The list of instance tags applied to all nodes.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	BootDiskKmsKey string `json:"bootDiskKmsKey,omitempty" yaml:"bootDiskKmsKey,omitempty"`

	// If enabled boot disks are configured with confidential mode.
	EnableConfidentialStorage bool `json:"enableConfidentialStorage,omitempty" yaml:"enableConfidentialStorage,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []Container_NodePoolNodeConfigGuestAccelerator `json:"guestAccelerators,omitempty" yaml:"guestAccelerators,omitempty"`

	// The name of a Google Compute Engine machine type.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// Enable or disable NCCL Fast Socket in the node pool.
	FastSocket Container_NodePoolNodeConfigFastSocket `json:"fastSocket,omitempty" yaml:"fastSocket,omitempty"`

	// The maintenance policy for the hosts on which the GKE VMs run on.
	HostMaintenancePolicy Container_NodePoolNodeConfigHostMaintenancePolicy `json:"hostMaintenancePolicy,omitempty" yaml:"hostMaintenancePolicy,omitempty"`

	// Parameters that can be configured on Linux nodes.
	LinuxNodeConfig Container_NodePoolNodeConfigLinuxNodeConfig `json:"linuxNodeConfig,omitempty" yaml:"linuxNodeConfig,omitempty"`

	// The Google Cloud Platform Service Account to be used by the node VMs.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Node affinity options for sole tenant node pools.
	SoleTenantConfig Container_NodePoolNodeConfigSoleTenantConfig `json:"soleTenantConfig,omitempty" yaml:"soleTenantConfig,omitempty"`

	// Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.
	EphemeralStorageConfig Container_NodePoolNodeConfigEphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty" yaml:"ephemeralStorageConfig,omitempty"`

	// Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.
	EphemeralStorageLocalSsdConfig Container_NodePoolNodeConfigEphemeralStorageLocalSsdConfig `json:"ephemeralStorageLocalSsdConfig,omitempty" yaml:"ephemeralStorageLocalSsdConfig,omitempty"`

	// The number of local SSD disks to be attached to the node.
	LocalSsdCount int `json:"localSsdCount,omitempty" yaml:"localSsdCount,omitempty"`

	// The reservation affinity configuration for the node pool.
	ReservationAffinity Container_NodePoolNodeConfigReservationAffinity `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`

	// Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.
	LoggingVariant string `json:"loggingVariant,omitempty" yaml:"loggingVariant,omitempty"`

	// The metadata key/value pairs assigned to instances in the cluster.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Whether the nodes are created as preemptible VM instances.
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`

	// Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on sole tenant nodes.
	NodeGroup string `json:"nodeGroup,omitempty" yaml:"nodeGroup,omitempty"`

	// Sandbox configuration for this node.
	SandboxConfig Container_NodePoolNodeConfigSandboxConfig `json:"sandboxConfig,omitempty" yaml:"sandboxConfig,omitempty"`

	// List of Kubernetes taints to be applied to each node.
	Taints []Container_NodePoolNodeConfigTaint `json:"taints,omitempty" yaml:"taints,omitempty"`

	// The image type to use for this node. Note that for a given image type, the latest version of it will be used.
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`

	// Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty" yaml:"resourceManagerTags,omitempty"`

	// The workload metadata configuration for this node.
	WorkloadMetadataConfig Container_NodePoolNodeConfigWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty" yaml:"workloadMetadataConfig,omitempty"`

	// GCFS configuration for this node.
	GcfsConfig Container_NodePoolNodeConfigGcfsConfig `json:"gcfsConfig,omitempty" yaml:"gcfsConfig,omitempty"`

	// Node kubelet configs.
	KubeletConfig Container_NodePoolNodeConfigKubeletConfig `json:"kubeletConfig,omitempty" yaml:"kubeletConfig,omitempty"`

	// Parameters for raw-block local NVMe SSDs.
	LocalNvmeSsdBlockConfig Container_NodePoolNodeConfigLocalNvmeSsdBlockConfig `json:"localNvmeSsdBlockConfig,omitempty" yaml:"localNvmeSsdBlockConfig,omitempty"`

	// Shielded Instance options.
	ShieldedInstanceConfig Container_NodePoolNodeConfigShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	// Specifies options for controlling advanced machine features.
	AdvancedMachineFeatures Container_NodePoolNodeConfigAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	// Configuration for Confidential Nodes feature. Structure is documented below.
	ConfidentialNodes Container_NodePoolNodeConfigConfidentialNodes `json:"confidentialNodes,omitempty" yaml:"confidentialNodes,omitempty"`

	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	// List of kubernetes taints applied to each node.
	EffectiveTaints []Container_NodePoolNodeConfigEffectiveTaint `json:"effectiveTaints,omitempty" yaml:"effectiveTaints,omitempty"`

	// Whether the nodes are created as spot VM instances.
	Spot bool `json:"spot,omitempty" yaml:"spot,omitempty"`

	// Enable or disable gvnic in the node pool.
	Gvnic Container_NodePoolNodeConfigGvnic `json:"gvnic,omitempty" yaml:"gvnic,omitempty"`

	// The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
