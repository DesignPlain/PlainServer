package types

type Eks_ClusterKubernetesNetworkConfig struct {
	// The IP family used to assign Kubernetes pod and service addresses. Valid values are `ipv4` (default) and `ipv6`. You can only specify an IP family when you create a cluster, changing this value will force a new cluster to be created.
	IpFamily string `json:"ipFamily,omitempty" yaml:"ipFamily,omitempty"`

	/*
	   The CIDR block to assign Kubernetes pod and service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. You can only specify a custom CIDR block when you create a cluster, changing this value will force a new cluster to be created. The block must meet the following requirements:

	   - Within one of the following private IP address blocks: 10.0.0.0/8, 172.16.0.0/12, or 192.168.0.0/16.

	   - Doesn't overlap with any CIDR block assigned to the VPC that you selected for VPC.

	   - Between /24 and /12.
	*/
	ServiceIpv4Cidr string `json:"serviceIpv4Cidr,omitempty" yaml:"serviceIpv4Cidr,omitempty"`

	// The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can't specify a custom IPv6 CIDR block when you create the cluster.
	ServiceIpv6Cidr string `json:"serviceIpv6Cidr,omitempty" yaml:"serviceIpv6Cidr,omitempty"`
}
