package types

type Fsx_getOntapFileSystemEndpoint struct {
	// A FileSystemEndpoint for managing your file system by setting up NetApp SnapMirror with other ONTAP systems. See FileSystemEndpoint below.
	Interclusters []Fsx_getOntapFileSystemEndpointIntercluster `json:"interclusters,omitempty" yaml:"interclusters,omitempty"`

	// A FileSystemEndpoint for managing your file system using the NetApp ONTAP CLI and NetApp ONTAP API. See FileSystemEndpoint below.
	Managements []Fsx_getOntapFileSystemEndpointManagement `json:"managements,omitempty" yaml:"managements,omitempty"`
}
