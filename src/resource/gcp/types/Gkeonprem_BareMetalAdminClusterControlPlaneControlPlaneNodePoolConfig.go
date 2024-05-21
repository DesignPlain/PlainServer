package types

type Gkeonprem_BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig struct {
	/*
	   The generic configuration for a node pool running the control plane.
	   Structure is documented below.
	*/
	NodePoolConfig Gkeonprem_BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig `json:"nodePoolConfig,omitempty" yaml:"nodePoolConfig,omitempty"`
}
