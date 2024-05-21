package types

type Gkeonprem_BareMetalClusterLoadBalancerMetalLbConfig struct {
	/*
	   Specifies the load balancer's node pool configuration.
	   Structure is documented below.
	*/
	LoadBalancerNodePoolConfig Gkeonprem_BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfig `json:"loadBalancerNodePoolConfig,omitempty" yaml:"loadBalancerNodePoolConfig,omitempty"`

	/*
	   AddressPools is a list of non-overlapping IP pools used by load balancer
	   typed services. All addresses must be routable to load balancer nodes.
	   IngressVIP must be included in the pools.
	   Structure is documented below.
	*/
	AddressPools []Gkeonprem_BareMetalClusterLoadBalancerMetalLbConfigAddressPool `json:"addressPools,omitempty" yaml:"addressPools,omitempty"`
}
