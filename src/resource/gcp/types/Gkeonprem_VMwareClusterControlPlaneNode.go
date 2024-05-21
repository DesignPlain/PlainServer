package types

type Gkeonprem_VMwareClusterControlPlaneNode struct {
	/*
	   AutoResizeConfig provides auto resizing configurations.
	   Structure is documented below.
	*/
	AutoResizeConfig Gkeonprem_VMwareClusterControlPlaneNodeAutoResizeConfig `json:"autoResizeConfig,omitempty" yaml:"autoResizeConfig,omitempty"`

	/*
	   The number of CPUs for each admin cluster node that serve as control planes
	   for this VMware User Cluster. (default: 4 CPUs)
	*/
	Cpus int `json:"cpus,omitempty" yaml:"cpus,omitempty"`

	/*
	   The megabytes of memory for each admin cluster node that serves as a
	   control plane for this VMware User Cluster (default: 8192 MB memory).
	*/
	Memory int `json:"memory,omitempty" yaml:"memory,omitempty"`

	/*
	   The number of control plane nodes for this VMware User Cluster.
	   (default: 1 replica).
	*/
	Replicas int `json:"replicas,omitempty" yaml:"replicas,omitempty"`

	/*
	   (Output)
	   Vsphere-specific config.
	   Structure is documented below.
	*/
	VsphereConfigs []Gkeonprem_VMwareClusterControlPlaneNodeVsphereConfig `json:"vsphereConfigs,omitempty" yaml:"vsphereConfigs,omitempty"`
}
