package types

type Eks_getClusterKubernetesNetworkConfig struct {
	// `ipv4` or `ipv6`.
	IpFamily string `json:"ipFamily,omitempty" yaml:"ipFamily,omitempty"`

	// The CIDR block to assign Kubernetes pod and service IP addresses from if `ipv4` was specified when the cluster was created.
	ServiceIpv4Cidr string `json:"serviceIpv4Cidr,omitempty" yaml:"serviceIpv4Cidr,omitempty"`

	// The CIDR block to assign Kubernetes pod and service IP addresses from if `ipv6` was specified when the cluster was created. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can't specify a custom IPv6 CIDR block when you create the cluster.
	ServiceIpv6Cidr string `json:"serviceIpv6Cidr,omitempty" yaml:"serviceIpv6Cidr,omitempty"`
}
