package types

type Datasync_LocationFsxOntapFileSystemProtocolNfs struct {
	// Mount options that are available for DataSync to access an NFS location. See NFS Mount Options below.
	MountOptions Datasync_LocationFsxOntapFileSystemProtocolNfsMountOptions `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`
}
