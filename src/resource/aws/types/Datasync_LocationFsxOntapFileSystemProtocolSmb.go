package types

type Datasync_LocationFsxOntapFileSystemProtocolSmb struct {
	// Fully qualified domain name of the Microsoft Active Directory (AD) that your storage virtual machine belongs to.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// Mount options that are available for DataSync to access an SMB location. See SMB Mount Options below.
	MountOptions Datasync_LocationFsxOntapFileSystemProtocolSmbMountOptions `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`

	// Password of a user who has permission to access your SVM.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Username that can mount the location and access the files, folders, and metadata that you need in the SVM.
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
