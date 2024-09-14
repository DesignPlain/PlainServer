package types

type Container_getClusterNodePoolNodeConfig struct {
	// Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.
	EphemeralStorageConfigs []Container_getClusterNodePoolNodeConfigEphemeralStorageConfig `json:"ephemeralStorageConfigs,omitempty" yaml:"ephemeralStorageConfigs,omitempty"`

	// The image type to use for this node. Note that for a given image type, the latest version of it will be used.
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`

	// Node kubelet configs.
	KubeletConfigs []Container_getClusterNodePoolNodeConfigKubeletConfig `json:"kubeletConfigs,omitempty" yaml:"kubeletConfigs,omitempty"`

	// Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	// The set of Google API scopes to be made available on all of the node VMs.
	OauthScopes []string `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`

	// The Google Cloud Platform Service Account to be used by the node VMs.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// List of Kubernetes taints to be applied to each node.
	Taints []Container_getClusterNodePoolNodeConfigTaint `json:"taints,omitempty" yaml:"taints,omitempty"`

	// Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after pool creation without deleting and recreating the entire pool.
	ConfidentialNodes []Container_getClusterNodePoolNodeConfigConfidentialNode `json:"confidentialNodes,omitempty" yaml:"confidentialNodes,omitempty"`

	// If enabled boot disks are configured with confidential mode.
	EnableConfidentialStorage bool `json:"enableConfidentialStorage,omitempty" yaml:"enableConfidentialStorage,omitempty"`

	// Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.
	EphemeralStorageLocalSsdConfigs []Container_getClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfig `json:"ephemeralStorageLocalSsdConfigs,omitempty" yaml:"ephemeralStorageLocalSsdConfigs,omitempty"`

	// Enable or disable gvnic in the node pool.
	Gvnics []Container_getClusterNodePoolNodeConfigGvnic `json:"gvnics,omitempty" yaml:"gvnics,omitempty"`

	// The name of a Google Compute Engine machine type.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// The metadata key/value pairs assigned to instances in the cluster.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	// The number of local SSD disks to be attached to the node.
	LocalSsdCount int `json:"localSsdCount,omitempty" yaml:"localSsdCount,omitempty"`

	// The GCE resource labels (a map of key/value pairs) to be applied to the node pool.
	ResourceLabels map[string]string `json:"resourceLabels,omitempty" yaml:"resourceLabels,omitempty"`

	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty" yaml:"resourceManagerTags,omitempty"`

	// The maintenance policy for the hosts on which the GKE VMs run on.
	HostMaintenancePolicies []Container_getClusterNodePoolNodeConfigHostMaintenancePolicy `json:"hostMaintenancePolicies,omitempty" yaml:"hostMaintenancePolicies,omitempty"`

	// The reservation affinity configuration for the node pool.
	ReservationAffinities []Container_getClusterNodePoolNodeConfigReservationAffinity `json:"reservationAffinities,omitempty" yaml:"reservationAffinities,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []Container_getClusterNodePoolNodeConfigGuestAccelerator `json:"guestAccelerators,omitempty" yaml:"guestAccelerators,omitempty"`

	// Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on sole tenant nodes.
	NodeGroup string `json:"nodeGroup,omitempty" yaml:"nodeGroup,omitempty"`

	// Whether the nodes are created as spot VM instances.
	Spot bool `json:"spot,omitempty" yaml:"spot,omitempty"`

	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`

	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	BootDiskKmsKey string `json:"bootDiskKmsKey,omitempty" yaml:"bootDiskKmsKey,omitempty"`

	// List of kubernetes taints applied to each node.
	EffectiveTaints []Container_getClusterNodePoolNodeConfigEffectiveTaint `json:"effectiveTaints,omitempty" yaml:"effectiveTaints,omitempty"`

	// Parameters for raw-block local NVMe SSDs.
	LocalNvmeSsdBlockConfigs []Container_getClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig `json:"localNvmeSsdBlockConfigs,omitempty" yaml:"localNvmeSsdBlockConfigs,omitempty"`

	// Specifies options for controlling advanced machine features.
	AdvancedMachineFeatures []Container_getClusterNodePoolNodeConfigAdvancedMachineFeature `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	// The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.
	LoggingVariant string `json:"loggingVariant,omitempty" yaml:"loggingVariant,omitempty"`

	// GCFS configuration for this node.
	GcfsConfigs []Container_getClusterNodePoolNodeConfigGcfsConfig `json:"gcfsConfigs,omitempty" yaml:"gcfsConfigs,omitempty"`

	// Parameters that can be configured on Linux nodes.
	LinuxNodeConfigs []Container_getClusterNodePoolNodeConfigLinuxNodeConfig `json:"linuxNodeConfigs,omitempty" yaml:"linuxNodeConfigs,omitempty"`

	// Whether the nodes are created as preemptible VM instances.
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	// Sandbox configuration for this node.
	SandboxConfigs []Container_getClusterNodePoolNodeConfigSandboxConfig `json:"sandboxConfigs,omitempty" yaml:"sandboxConfigs,omitempty"`

	// Shielded Instance options.
	ShieldedInstanceConfigs []Container_getClusterNodePoolNodeConfigShieldedInstanceConfig `json:"shieldedInstanceConfigs,omitempty" yaml:"shieldedInstanceConfigs,omitempty"`

	// Node affinity options for sole tenant node pools.
	SoleTenantConfigs []Container_getClusterNodePoolNodeConfigSoleTenantConfig `json:"soleTenantConfigs,omitempty" yaml:"soleTenantConfigs,omitempty"`

	// The list of instance tags applied to all nodes.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The workload metadata configuration for this node.
	WorkloadMetadataConfigs []Container_getClusterNodePoolNodeConfigWorkloadMetadataConfig `json:"workloadMetadataConfigs,omitempty" yaml:"workloadMetadataConfigs,omitempty"`

	// Enable or disable NCCL Fast Socket in the node pool.
	FastSockets []Container_getClusterNodePoolNodeConfigFastSocket `json:"fastSockets,omitempty" yaml:"fastSockets,omitempty"`
}
