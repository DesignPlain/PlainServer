package types

type Container_getClusterNodeConfig struct {
	// Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.
	EphemeralStorageConfigs []Container_getClusterNodeConfigEphemeralStorageConfig `json:"ephemeralStorageConfigs,omitempty" yaml:"ephemeralStorageConfigs,omitempty"`

	// Whether the nodes are created as preemptible VM instances.
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	BootDiskKmsKey string `json:"bootDiskKmsKey,omitempty" yaml:"bootDiskKmsKey,omitempty"`

	// The image type to use for this node. Note that for a given image type, the latest version of it will be used.
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`

	// The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of a Google Compute Engine machine type.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	// Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on sole tenant nodes.
	NodeGroup string `json:"nodeGroup,omitempty" yaml:"nodeGroup,omitempty"`

	// GCFS configuration for this node.
	GcfsConfigs []Container_getClusterNodeConfigGcfsConfig `json:"gcfsConfigs,omitempty" yaml:"gcfsConfigs,omitempty"`

	// Enable or disable NCCL Fast Socket in the node pool.
	FastSockets []Container_getClusterNodeConfigFastSocket `json:"fastSockets,omitempty" yaml:"fastSockets,omitempty"`

	// The number of local SSD disks to be attached to the node.
	LocalSsdCount int `json:"localSsdCount,omitempty" yaml:"localSsdCount,omitempty"`

	// The reservation affinity configuration for the node pool.
	ReservationAffinities []Container_getClusterNodeConfigReservationAffinity `json:"reservationAffinities,omitempty" yaml:"reservationAffinities,omitempty"`

	// The GCE resource labels (a map of key/value pairs) to be applied to the node pool.
	ResourceLabels map[string]string `json:"resourceLabels,omitempty" yaml:"resourceLabels,omitempty"`

	// Shielded Instance options.
	ShieldedInstanceConfigs []Container_getClusterNodeConfigShieldedInstanceConfig `json:"shieldedInstanceConfigs,omitempty" yaml:"shieldedInstanceConfigs,omitempty"`

	// Specifies options for controlling advanced machine features.
	AdvancedMachineFeatures []Container_getClusterNodeConfigAdvancedMachineFeature `json:"advancedMachineFeatures,omitempty" yaml:"advancedMachineFeatures,omitempty"`

	// List of kubernetes taints applied to each node.
	EffectiveTaints []Container_getClusterNodeConfigEffectiveTaint `json:"effectiveTaints,omitempty" yaml:"effectiveTaints,omitempty"`

	// Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk.
	EphemeralStorageLocalSsdConfigs []Container_getClusterNodeConfigEphemeralStorageLocalSsdConfig `json:"ephemeralStorageLocalSsdConfigs,omitempty" yaml:"ephemeralStorageLocalSsdConfigs,omitempty"`

	// Enable or disable gvnic in the node pool.
	Gvnics []Container_getClusterNodeConfigGvnic `json:"gvnics,omitempty" yaml:"gvnics,omitempty"`

	// Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.
	LoggingVariant string `json:"loggingVariant,omitempty" yaml:"loggingVariant,omitempty"`

	// The metadata key/value pairs assigned to instances in the cluster.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty" yaml:"resourceManagerTags,omitempty"`

	// The Google Cloud Platform Service Account to be used by the node VMs.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	// Whether the nodes are created as spot VM instances.
	Spot bool `json:"spot,omitempty" yaml:"spot,omitempty"`

	// Node affinity options for sole tenant node pools.
	SoleTenantConfigs []Container_getClusterNodeConfigSoleTenantConfig `json:"soleTenantConfigs,omitempty" yaml:"soleTenantConfigs,omitempty"`

	// The maintenance policy for the hosts on which the GKE VMs run on.
	HostMaintenancePolicies []Container_getClusterNodeConfigHostMaintenancePolicy `json:"hostMaintenancePolicies,omitempty" yaml:"hostMaintenancePolicies,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []Container_getClusterNodeConfigGuestAccelerator `json:"guestAccelerators,omitempty" yaml:"guestAccelerators,omitempty"`

	// Parameters for raw-block local NVMe SSDs.
	LocalNvmeSsdBlockConfigs []Container_getClusterNodeConfigLocalNvmeSsdBlockConfig `json:"localNvmeSsdBlockConfigs,omitempty" yaml:"localNvmeSsdBlockConfigs,omitempty"`

	// The set of Google API scopes to be made available on all of the node VMs.
	OauthScopes []string `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`

	// The list of instance tags applied to all nodes.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// List of Kubernetes taints to be applied to each node.
	Taints []Container_getClusterNodeConfigTaint `json:"taints,omitempty" yaml:"taints,omitempty"`

	// The workload metadata configuration for this node.
	WorkloadMetadataConfigs []Container_getClusterNodeConfigWorkloadMetadataConfig `json:"workloadMetadataConfigs,omitempty" yaml:"workloadMetadataConfigs,omitempty"`

	// If enabled boot disks are configured with confidential mode.
	EnableConfidentialStorage bool `json:"enableConfidentialStorage,omitempty" yaml:"enableConfidentialStorage,omitempty"`

	// Sandbox configuration for this node.
	SandboxConfigs []Container_getClusterNodeConfigSandboxConfig `json:"sandboxConfigs,omitempty" yaml:"sandboxConfigs,omitempty"`

	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`

	// Node kubelet configs.
	KubeletConfigs []Container_getClusterNodeConfigKubeletConfig `json:"kubeletConfigs,omitempty" yaml:"kubeletConfigs,omitempty"`

	// Parameters that can be configured on Linux nodes.
	LinuxNodeConfigs []Container_getClusterNodeConfigLinuxNodeConfig `json:"linuxNodeConfigs,omitempty" yaml:"linuxNodeConfigs,omitempty"`

	// Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after pool creation without deleting and recreating the entire pool.
	ConfidentialNodes []Container_getClusterNodeConfigConfidentialNode `json:"confidentialNodes,omitempty" yaml:"confidentialNodes,omitempty"`
}
