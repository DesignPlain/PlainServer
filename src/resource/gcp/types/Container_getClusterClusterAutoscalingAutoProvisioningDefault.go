package types

type Container_getClusterClusterAutoscalingAutoProvisioningDefault struct {
	// The default image type used by NAP once a new node pool is being created.
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`

	// Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell.
	MinCpuPlatform string `json:"minCpuPlatform,omitempty" yaml:"minCpuPlatform,omitempty"`

	// Shielded Instance options.
	ShieldedInstanceConfigs []Container_getClusterClusterAutoscalingAutoProvisioningDefaultShieldedInstanceConfig `json:"shieldedInstanceConfigs,omitempty" yaml:"shieldedInstanceConfigs,omitempty"`

	// Specifies the upgrade settings for NAP created node pools
	UpgradeSettings []Container_getClusterClusterAutoscalingAutoProvisioningDefaultUpgradeSetting `json:"upgradeSettings,omitempty" yaml:"upgradeSettings,omitempty"`

	// Scopes that are used by NAP when creating node pools.
	OauthScopes []string `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`

	// The Google Cloud Platform Service Account to be used by the node VMs.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	BootDiskKmsKey string `json:"bootDiskKmsKey,omitempty" yaml:"bootDiskKmsKey,omitempty"`

	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	DiskSize int `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`

	// Type of the disk attached to each node.
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`

	// NodeManagement configuration for this NodePool.
	Managements []Container_getClusterClusterAutoscalingAutoProvisioningDefaultManagement `json:"managements,omitempty" yaml:"managements,omitempty"`
}
