package types

type Gkeonprem_BareMetalAdminClusterLoadBalancer struct {
	/*
	   A nested object resource
	   Structure is documented below.
	*/
	ManualLbConfig Gkeonprem_BareMetalAdminClusterLoadBalancerManualLbConfig `json:"manualLbConfig,omitempty" yaml:"manualLbConfig,omitempty"`

	/*
	   Specifies the load balancer ports.
	   Structure is documented below.
	*/
	PortConfig Gkeonprem_BareMetalAdminClusterLoadBalancerPortConfig `json:"portConfig,omitempty" yaml:"portConfig,omitempty"`

	/*
	   Specified the Bare Metal Load Balancer Config
	   Structure is documented below.
	*/
	VipConfig Gkeonprem_BareMetalAdminClusterLoadBalancerVipConfig `json:"vipConfig,omitempty" yaml:"vipConfig,omitempty"`
}
