package types

type Datasync_FsxOpenZfsFileSystemProtocol struct {
	// Represents the Network File System (NFS) protocol that DataSync uses to access your FSx for OpenZFS file system. See below.
	Nfs Datasync_FsxOpenZfsFileSystemProtocolNfs `json:"nfs,omitempty" yaml:"nfs,omitempty"`
}
