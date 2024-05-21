package types

type Container_AzureClusterControlPlane struct {
	// Optional. Configuration related to the root volume provisioned for each control plane replica. When unspecified, it defaults to 32-GiB Azure Disk.
	RootVolume Container_AzureClusterControlPlaneRootVolume `json:"rootVolume,omitempty" yaml:"rootVolume,omitempty"`

	// The Kubernetes version to run on control plane replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAzureServerConfig.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. When unspecified, it defaults to a 8-GiB Azure Disk.
	MainVolume Container_AzureClusterControlPlaneMainVolume `json:"mainVolume,omitempty" yaml:"mainVolume,omitempty"`

	// Proxy configuration for outbound HTTP(S) traffic.
	ProxyConfig Container_AzureClusterControlPlaneProxyConfig `json:"proxyConfig,omitempty" yaml:"proxyConfig,omitempty"`

	// SSH configuration for how to access the underlying control plane machines.
	SshConfig Container_AzureClusterControlPlaneSshConfig `json:"sshConfig,omitempty" yaml:"sshConfig,omitempty"`

	// The ARM ID of the subnet where the control plane VMs are deployed. Example: `/subscriptions//resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/default`.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// Optional. A set of tags to apply to all underlying control plane Azure resources.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`. For available VM sizes, see https://docs.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions. When unspecified, it defaults to `Standard_DS2_v2`.
	VmSize string `json:"vmSize,omitempty" yaml:"vmSize,omitempty"`

	// Optional. Configuration related to application-layer secrets encryption.
	DatabaseEncryption Container_AzureClusterControlPlaneDatabaseEncryption `json:"databaseEncryption,omitempty" yaml:"databaseEncryption,omitempty"`

	// Configuration for where to place the control plane replicas. Up to three replica placement instances can be specified. If replica_placements is set, the replica placement instances will be applied to the three control plane replicas as evenly as possible.
	ReplicaPlacements []Container_AzureClusterControlPlaneReplicaPlacement `json:"replicaPlacements,omitempty" yaml:"replicaPlacements,omitempty"`
}
