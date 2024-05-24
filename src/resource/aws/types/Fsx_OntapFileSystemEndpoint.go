package types

type Fsx_OntapFileSystemEndpoint struct {
	// An endpoint for managing your file system by setting up NetApp SnapMirror with other ONTAP systems. See Endpoint.
	Interclusters []Fsx_OntapFileSystemEndpointIntercluster `json:"interclusters,omitempty" yaml:"interclusters,omitempty"`

	// An endpoint for managing your file system using the NetApp ONTAP CLI and NetApp ONTAP API. See Endpoint.
	Managements []Fsx_OntapFileSystemEndpointManagement `json:"managements,omitempty" yaml:"managements,omitempty"`
}
