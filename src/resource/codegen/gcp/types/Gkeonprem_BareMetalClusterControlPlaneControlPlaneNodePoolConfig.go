package types

type Gkeonprem_BareMetalClusterControlPlaneControlPlaneNodePoolConfig struct {
	/*
	   The generic configuration for a node pool running the control plane.
	   Structure is documented below.
	*/
	NodePoolConfig Gkeonprem_BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig `json:"nodePoolConfig,omitempty" yaml:"nodePoolConfig,omitempty"`
}
