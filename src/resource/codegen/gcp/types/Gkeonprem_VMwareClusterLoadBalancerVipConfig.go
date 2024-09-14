package types

type Gkeonprem_VMwareClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this cluster.
	ControlPlaneVip string `json:"controlPlaneVip,omitempty" yaml:"controlPlaneVip,omitempty"`

	/*
	   The VIP which you previously set aside for ingress traffic into this cluster.

	   <a name="nested_f5_config"></a>The `f5_config` block supports:
	*/
	IngressVip string `json:"ingressVip,omitempty" yaml:"ingressVip,omitempty"`
}
