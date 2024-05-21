package types

type Gkeonprem_VMwareClusterControlPlaneNodeVsphereConfig struct {
	/*
	   (Output)
	   The Vsphere datastore used by the Control Plane Node.
	*/
	Datastore string `json:"datastore,omitempty" yaml:"datastore,omitempty"`

	/*
	   (Output)
	   The Vsphere storage policy used by the control plane Node.

	   - - -
	*/
	StoragePolicyName string `json:"storagePolicyName,omitempty" yaml:"storagePolicyName,omitempty"`
}
