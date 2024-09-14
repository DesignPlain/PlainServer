package types

type Gkeonprem_VMwareClusterControlPlaneNodeAutoResizeConfig struct {
	/*
	   Whether to enable control plane node auto resizing.

	   <a name="nested_vsphere_config"></a>The `vsphere_config` block contains:
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
