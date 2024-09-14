package types

type Fsx_OntapStorageVirtualMachineEndpoint struct {
	// An endpoint for accessing data on your storage virtual machine via iSCSI protocol. See Endpoint.
	Iscsis []Fsx_OntapStorageVirtualMachineEndpointIscsi `json:"iscsis,omitempty" yaml:"iscsis,omitempty"`

	// An endpoint for managing your file system using the NetApp ONTAP CLI and NetApp ONTAP API. See Endpoint.
	Managements []Fsx_OntapStorageVirtualMachineEndpointManagement `json:"managements,omitempty" yaml:"managements,omitempty"`

	// An endpoint for accessing data on your storage virtual machine via NFS protocol. See Endpoint.
	Nfs []Fsx_OntapStorageVirtualMachineEndpointNf `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	// An endpoint for accessing data on your storage virtual machine via SMB protocol. This is only set if an active_directory_configuration has been set. See Endpoint.
	Smbs []Fsx_OntapStorageVirtualMachineEndpointSmb `json:"smbs,omitempty" yaml:"smbs,omitempty"`
}
