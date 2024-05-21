package types

type Gkeonprem_BareMetalClusterLoadBalancerBgpLbConfig struct {
	/*
	   BGP autonomous system number (ASN) of the cluster.
	   This field can be updated after cluster creation.
	*/
	Asn int `json:"asn,omitempty" yaml:"asn,omitempty"`

	/*
	   The list of BGP peers that the cluster will connect to.
	   At least one peer must be configured for each control plane node.
	   Control plane nodes will connect to these peers to advertise the control
	   plane VIP. The Services load balancer also uses these peers by default.
	   This field can be updated after cluster creation.
	   Structure is documented below.
	*/
	BgpPeerConfigs []Gkeonprem_BareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfig `json:"bgpPeerConfigs,omitempty" yaml:"bgpPeerConfigs,omitempty"`

	/*
	   Specifies the node pool running data plane load balancing. L2 connectivity
	   is required among nodes in this pool. If missing, the control plane node
	   pool is used for data plane load balancing.
	   Structure is documented below.
	*/
	LoadBalancerNodePoolConfig Gkeonprem_BareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig `json:"loadBalancerNodePoolConfig,omitempty" yaml:"loadBalancerNodePoolConfig,omitempty"`

	/*
	   AddressPools is a list of non-overlapping IP pools used by load balancer
	   typed services. All addresses must be routable to load balancer nodes.
	   IngressVIP must be included in the pools.
	   Structure is documented below.
	*/
	AddressPools []Gkeonprem_BareMetalClusterLoadBalancerBgpLbConfigAddressPool `json:"addressPools,omitempty" yaml:"addressPools,omitempty"`
}
