package types

type Gkeonprem_VMwareClusterLoadBalancer struct {
	/*
	   Configuration for F5 Big IP typed load balancers.
	   Structure is documented below.
	*/
	F5Config Gkeonprem_VMwareClusterLoadBalancerF5Config `json:"f5Config,omitempty" yaml:"f5Config,omitempty"`

	/*
	   Manually configured load balancers.
	   Structure is documented below.
	*/
	ManualLbConfig Gkeonprem_VMwareClusterLoadBalancerManualLbConfig `json:"manualLbConfig,omitempty" yaml:"manualLbConfig,omitempty"`

	/*
	   Configuration for MetalLB typed load balancers.
	   Structure is documented below.
	*/
	MetalLbConfig Gkeonprem_VMwareClusterLoadBalancerMetalLbConfig `json:"metalLbConfig,omitempty" yaml:"metalLbConfig,omitempty"`

	/*
	   The VIPs used by the load balancer.
	   Structure is documented below.
	*/
	VipConfig Gkeonprem_VMwareClusterLoadBalancerVipConfig `json:"vipConfig,omitempty" yaml:"vipConfig,omitempty"`
}
