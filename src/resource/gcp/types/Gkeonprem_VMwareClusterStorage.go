package types

type Gkeonprem_VMwareClusterStorage struct {
	/*
	   Whether or not to deploy vSphere CSI components in the VMware User Cluster.
	   Enabled by default.
	*/
	VsphereCsiDisabled bool `json:"vsphereCsiDisabled,omitempty" yaml:"vsphereCsiDisabled,omitempty"`
}
