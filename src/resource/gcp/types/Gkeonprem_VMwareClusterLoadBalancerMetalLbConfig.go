package types

type Gkeonprem_VMwareClusterLoadBalancerMetalLbConfig struct {
	/*
	   AddressPools is a list of non-overlapping IP pools used by load balancer
	   typed services. All addresses must be routable to load balancer nodes.
	   IngressVIP must be included in the pools.
	   Structure is documented below.
	*/
	AddressPools []Gkeonprem_VMwareClusterLoadBalancerMetalLbConfigAddressPool `json:"addressPools,omitempty" yaml:"addressPools,omitempty"`
}
