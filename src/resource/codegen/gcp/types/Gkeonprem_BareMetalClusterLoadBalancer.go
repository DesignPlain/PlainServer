package types

type Gkeonprem_BareMetalClusterLoadBalancer struct {
	/*
	   Configuration for BGP typed load balancers.
	   Structure is documented below.
	*/
	BgpLbConfig Gkeonprem_BareMetalClusterLoadBalancerBgpLbConfig `json:"bgpLbConfig,omitempty" yaml:"bgpLbConfig,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	ManualLbConfig Gkeonprem_BareMetalClusterLoadBalancerManualLbConfig `json:"manualLbConfig,omitempty" yaml:"manualLbConfig,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	MetalLbConfig Gkeonprem_BareMetalClusterLoadBalancerMetalLbConfig `json:"metalLbConfig,omitempty" yaml:"metalLbConfig,omitempty"`

	/*
	   Specifies the load balancer ports.
	   Structure is documented below.
	*/
	PortConfig Gkeonprem_BareMetalClusterLoadBalancerPortConfig `json:"portConfig,omitempty" yaml:"portConfig,omitempty"`

	/*
	   Specified the Bare Metal Load Balancer Config
	   Structure is documented below.
	*/
	VipConfig Gkeonprem_BareMetalClusterLoadBalancerVipConfig `json:"vipConfig,omitempty" yaml:"vipConfig,omitempty"`
}
