package types

type Container_AzureNodePoolConfig struct {
	// Optional. A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to `Standard_DS2_v2`.
	VmSize string `json:"vmSize,omitempty" yaml:"vmSize,omitempty"`

	// The OS image type to use on node pool instances.
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`

	// Optional. The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Proxy configuration for outbound HTTP(S) traffic.
	ProxyConfig Container_AzureNodePoolConfigProxyConfig `json:"proxyConfig,omitempty" yaml:"proxyConfig,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each node pool machine. When unspecified, it defaults to a 32-GiB Azure Disk.
	RootVolume Container_AzureNodePoolConfigRootVolume `json:"rootVolume,omitempty" yaml:"rootVolume,omitempty"`

	// SSH configuration for how to access the node pool machines.
	SshConfig Container_AzureNodePoolConfigSshConfig `json:"sshConfig,omitempty" yaml:"sshConfig,omitempty"`
}
