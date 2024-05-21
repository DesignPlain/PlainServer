package types

type Container_ClusterIpAllocationPolicy struct {
	/*
	   The IP address range for the cluster pod IPs.
	   Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14)
	   to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14)
	   from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to
	   pick a specific range to use.
	*/
	ClusterIpv4CidrBlock string `json:"clusterIpv4CidrBlock,omitempty" yaml:"clusterIpv4CidrBlock,omitempty"`

	/*
	   The name of the existing secondary
	   range in the cluster's subnetwork to use for pod IP addresses. Alternatively,
	   `cluster_ipv4_cidr_block` can be used to automatically create a GKE-managed one.
	*/
	ClusterSecondaryRangeName string `json:"clusterSecondaryRangeName,omitempty" yaml:"clusterSecondaryRangeName,omitempty"`

	// Configuration for cluster level pod cidr overprovision. Default is disabled=false.
	PodCidrOverprovisionConfig Container_ClusterIpAllocationPolicyPodCidrOverprovisionConfig `json:"podCidrOverprovisionConfig,omitempty" yaml:"podCidrOverprovisionConfig,omitempty"`

	/*
	   The IP address range of the services IPs in this cluster.
	   Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14)
	   to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14)
	   from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to
	   pick a specific range to use.
	*/
	ServicesIpv4CidrBlock string `json:"servicesIpv4CidrBlock,omitempty" yaml:"servicesIpv4CidrBlock,omitempty"`

	/*
	   The name of the existing
	   secondary range in the cluster's subnetwork to use for service `ClusterIP`s.
	   Alternatively, `services_ipv4_cidr_block` can be used to automatically create a
	   GKE-managed one.
	*/
	ServicesSecondaryRangeName string `json:"servicesSecondaryRangeName,omitempty" yaml:"servicesSecondaryRangeName,omitempty"`

	/*
	   The IP Stack Type of the cluster.
	   Default value is `IPV4`.
	   Possible values are `IPV4` and `IPV4_IPV6`.
	*/
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`

	/*
	   The configuration for additional pod secondary ranges at
	   the cluster level. Used for Autopilot clusters and Standard clusters with which control of the
	   secondary Pod IP address assignment to node pools isn't needed. Structure is documented below.
	*/
	AdditionalPodRangesConfig Container_ClusterIpAllocationPolicyAdditionalPodRangesConfig `json:"additionalPodRangesConfig,omitempty" yaml:"additionalPodRangesConfig,omitempty"`
}
