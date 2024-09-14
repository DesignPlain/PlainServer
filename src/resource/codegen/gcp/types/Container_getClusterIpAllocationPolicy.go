package types

type Container_getClusterIpAllocationPolicy struct {
	// The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	ClusterIpv4CidrBlock string `json:"clusterIpv4CidrBlock,omitempty" yaml:"clusterIpv4CidrBlock,omitempty"`

	// The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_ipv4_cidr_block can be used to automatically create a GKE-managed one.
	ClusterSecondaryRangeName string `json:"clusterSecondaryRangeName,omitempty" yaml:"clusterSecondaryRangeName,omitempty"`

	// Configuration for cluster level pod cidr overprovision. Default is disabled=false.
	PodCidrOverprovisionConfigs []Container_getClusterIpAllocationPolicyPodCidrOverprovisionConfig `json:"podCidrOverprovisionConfigs,omitempty" yaml:"podCidrOverprovisionConfigs,omitempty"`

	// The IP address range of the services IPs in this cluster. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	ServicesIpv4CidrBlock string `json:"servicesIpv4CidrBlock,omitempty" yaml:"servicesIpv4CidrBlock,omitempty"`

	// The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Alternatively, services_ipv4_cidr_block can be used to automatically create a GKE-managed one.
	ServicesSecondaryRangeName string `json:"servicesSecondaryRangeName,omitempty" yaml:"servicesSecondaryRangeName,omitempty"`

	// The IP Stack type of the cluster. Choose between IPV4 and IPV4_IPV6. Default type is IPV4 Only if not set
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`

	// AdditionalPodRangesConfig is the configuration for additional pod secondary ranges supporting the ClusterUpdate message.
	AdditionalPodRangesConfigs []Container_getClusterIpAllocationPolicyAdditionalPodRangesConfig `json:"additionalPodRangesConfigs,omitempty" yaml:"additionalPodRangesConfigs,omitempty"`
}
