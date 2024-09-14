package fsx

import types "libds/aws/types"

type OntapStorageVirtualMachine struct {
	// Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
	RootVolumeSecurityStyle string `json:"rootVolumeSecurityStyle,omitempty" yaml:"rootVolumeSecurityStyle,omitempty"`

	// Specifies the password to use when logging on to the SVM using a secure shell (SSH) connection to the SVM's management endpoint. Doing so enables you to manage the SVM using the NetApp ONTAP CLI or REST API. If you do not specify a password, you can still use the file system's fsxadmin user to manage the SVM.
	SvmAdminPassword string `json:"svmAdminPassword,omitempty" yaml:"svmAdminPassword,omitempty"`

	// A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
	ActiveDirectoryConfiguration types.Fsx_OntapStorageVirtualMachineActiveDirectoryConfiguration `json:"activeDirectoryConfiguration,omitempty" yaml:"activeDirectoryConfiguration,omitempty"`

	// The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
