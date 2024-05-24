package types

type Datasync_LocationFsxOntapFileSystemProtocolSmbMountOptions struct {
	// The specific NFS version that you want DataSync to use for mounting your NFS share. Valid values: `NFS3`. Default: `NFS3`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
