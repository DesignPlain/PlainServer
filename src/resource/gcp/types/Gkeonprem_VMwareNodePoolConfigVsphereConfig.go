package types

type Gkeonprem_VMwareNodePoolConfigVsphereConfig struct {
	// The name of the vCenter datastore. Inherited from the user cluster.
	Datastore string `json:"datastore,omitempty" yaml:"datastore,omitempty"`

	// Vsphere host groups to apply to all VMs in the node pool
	HostGroups []string `json:"hostGroups,omitempty" yaml:"hostGroups,omitempty"`

	/*
	   Tags to apply to VMs.
	   Structure is documented below.
	*/
	Tags []Gkeonprem_VMwareNodePoolConfigVsphereConfigTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
