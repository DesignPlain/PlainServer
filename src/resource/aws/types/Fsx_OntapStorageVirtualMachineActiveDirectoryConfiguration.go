package types

type Fsx_OntapStorageVirtualMachineActiveDirectoryConfiguration struct {
	// The NetBIOS name of the Active Directory computer object that will be created for your SVM. This is often the same as the SVM name but can be different. AWS limits to 15 characters because of standard NetBIOS naming limits.
	NetbiosName string `json:"netbiosName,omitempty" yaml:"netbiosName,omitempty"`

	//
	SelfManagedActiveDirectoryConfiguration Fsx_OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration `json:"selfManagedActiveDirectoryConfiguration,omitempty" yaml:"selfManagedActiveDirectoryConfiguration,omitempty"`
}
