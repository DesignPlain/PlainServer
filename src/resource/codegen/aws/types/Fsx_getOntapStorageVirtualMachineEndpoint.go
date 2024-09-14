package types

type Fsx_getOntapStorageVirtualMachineEndpoint struct {
	//
	Iscsis []Fsx_getOntapStorageVirtualMachineEndpointIscsi `json:"iscsis,omitempty" yaml:"iscsis,omitempty"`

	// An endpoint for managing SVMs using the NetApp ONTAP CLI, NetApp ONTAP API, or NetApp CloudManager. See SVM Endpoint below.
	Managements []Fsx_getOntapStorageVirtualMachineEndpointManagement `json:"managements,omitempty" yaml:"managements,omitempty"`

	// An endpoint for connecting using the Network File System (NFS) protocol. See SVM Endpoint below.
	Nfs []Fsx_getOntapStorageVirtualMachineEndpointNf `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	// An endpoint for connecting using the Server Message Block (SMB) protocol. See SVM Endpoint below.
	Smbs []Fsx_getOntapStorageVirtualMachineEndpointSmb `json:"smbs,omitempty" yaml:"smbs,omitempty"`
}
