package types

type Datasync_LocationFsxOntapFileSystemProtocol struct {
	// Network File System (NFS) protocol that DataSync uses to access your FSx ONTAP file system. See NFS below.
	Nfs Datasync_LocationFsxOntapFileSystemProtocolNfs `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	// Server Message Block (SMB) protocol that DataSync uses to access your FSx ONTAP file system. See [SMB] (#smb) below.
	Smb Datasync_LocationFsxOntapFileSystemProtocolSmb `json:"smb,omitempty" yaml:"smb,omitempty"`
}
