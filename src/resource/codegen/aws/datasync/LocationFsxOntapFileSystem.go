package datasync

import types "libds/aws/types"

type LocationFsxOntapFileSystem struct {
	// The data transfer protocol that DataSync uses to access your Amazon FSx file system. See Protocol below.
	Protocol types.Datasync_LocationFsxOntapFileSystemProtocol `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	/*
	   The security groups that provide access to your file system's preferred subnet. The security groups must allow outbbound traffic on the following ports (depending on the protocol you use):
	   - Network File System (NFS): TCP ports 111, 635, and 2049
	   - Server Message Block (SMB): TCP port 445
	*/
	SecurityGroupArns []string `json:"securityGroupArns,omitempty" yaml:"securityGroupArns,omitempty"`

	/*
	   The ARN of the SVM in your file system where you want to copy data to of from.

	   The following arguments are optional:
	*/
	StorageVirtualMachineArn string `json:"storageVirtualMachineArn,omitempty" yaml:"storageVirtualMachineArn,omitempty"`

	// Path to the file share in the SVM where you'll copy your data. You can specify a junction path (also known as a mount point), qtree path (for NFS file shares), or share name (for SMB file shares) (e.g. `/vol1`, `/vol1/tree1`, `share1`).
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
