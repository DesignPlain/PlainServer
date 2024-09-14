package types

type Gkeonprem_BareMetalClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for ingress traffic into this Bare Metal User Cluster.
	IngressVip string `json:"ingressVip,omitempty" yaml:"ingressVip,omitempty"`

	// The VIP which you previously set aside for the Kubernetes API of this Bare Metal User Cluster.
	ControlPlaneVip string `json:"controlPlaneVip,omitempty" yaml:"controlPlaneVip,omitempty"`
}
