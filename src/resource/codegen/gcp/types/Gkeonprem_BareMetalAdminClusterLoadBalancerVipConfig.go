package types

type Gkeonprem_BareMetalAdminClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this Bare Metal Admin Cluster.
	ControlPlaneVip string `json:"controlPlaneVip,omitempty" yaml:"controlPlaneVip,omitempty"`
}
