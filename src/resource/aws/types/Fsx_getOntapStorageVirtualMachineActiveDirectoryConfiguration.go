package types

type Fsx_getOntapStorageVirtualMachineActiveDirectoryConfiguration struct {
	// The NetBIOS name of the AD computer object to which the SVM is joined.
	NetbiosName string `json:"netbiosName,omitempty" yaml:"netbiosName,omitempty"`

	//
	SelfManagedActiveDirectoryConfigurations []Fsx_getOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration `json:"selfManagedActiveDirectoryConfigurations,omitempty" yaml:"selfManagedActiveDirectoryConfigurations,omitempty"`
}
