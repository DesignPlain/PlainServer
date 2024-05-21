package types

type Gkeonprem_VMwareClusterVcenter struct {
	// The name of the vCenter folder for the user cluster.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	// The name of the vCenter resource pool for the user cluster.
	ResourcePool string `json:"resourcePool,omitempty" yaml:"resourcePool,omitempty"`

	// The name of the vCenter storage policy for the user cluster.
	StoragePolicyName string `json:"storagePolicyName,omitempty" yaml:"storagePolicyName,omitempty"`

	/*
	   (Output)
	   The vCenter IP address.
	*/
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Contains the vCenter CA certificate public key for SSL verification.
	CaCertData string `json:"caCertData,omitempty" yaml:"caCertData,omitempty"`

	// The name of the vCenter cluster for the user cluster.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// The name of the vCenter datacenter for the user cluster.
	Datacenter string `json:"datacenter,omitempty" yaml:"datacenter,omitempty"`

	// The name of the vCenter datastore for the user cluster.
	Datastore string `json:"datastore,omitempty" yaml:"datastore,omitempty"`
}
