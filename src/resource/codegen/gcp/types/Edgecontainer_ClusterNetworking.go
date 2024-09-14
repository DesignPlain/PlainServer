package types

type Edgecontainer_ClusterNetworking struct {
	/*
	   All services in the cluster are assigned an RFC1918 IPv4 address from these
	   blocks. Only a single block is supported. This field cannot be changed
	   after creation.
	*/
	ServicesIpv4CidrBlocks []string `json:"servicesIpv4CidrBlocks,omitempty" yaml:"servicesIpv4CidrBlocks,omitempty"`

	/*
	   If specified, dual stack mode is enabled and all services in the cluster are
	   assigned an IPv6 address from these blocks alongside from an IPv4
	   address. Only a single block is supported. This field cannot be changed
	   after creation.
	*/
	ServicesIpv6CidrBlocks []string `json:"servicesIpv6CidrBlocks,omitempty" yaml:"servicesIpv6CidrBlocks,omitempty"`

	/*
	   All pods in the cluster are assigned an RFC1918 IPv4 address from these
	   blocks. Only a single block is supported. This field cannot be changed
	   after creation.
	*/
	ClusterIpv4CidrBlocks []string `json:"clusterIpv4CidrBlocks,omitempty" yaml:"clusterIpv4CidrBlocks,omitempty"`

	/*
	   If specified, dual stack mode is enabled and all pods in the cluster are
	   assigned an IPv6 address from these blocks alongside from an IPv4
	   address. Only a single block is supported. This field cannot be changed
	   after creation.
	*/
	ClusterIpv6CidrBlocks []string `json:"clusterIpv6CidrBlocks,omitempty" yaml:"clusterIpv6CidrBlocks,omitempty"`

	/*
	   (Output)
	   IP addressing type of this cluster i.e. SINGLESTACK_V4 vs DUALSTACK_V4_V6.
	*/
	NetworkType string `json:"networkType,omitempty" yaml:"networkType,omitempty"`
}
