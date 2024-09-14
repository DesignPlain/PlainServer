package types

type Licensemanager_getReceivedLicenseConsumptionConfiguration struct {
	// Details about a borrow configuration. Detailed below
	BorrowConfigurations []Licensemanager_getReceivedLicenseConsumptionConfigurationBorrowConfiguration `json:"borrowConfigurations,omitempty" yaml:"borrowConfigurations,omitempty"`

	// Details about a provisional configuration. Detailed below
	ProvisionalConfigurations []Licensemanager_getReceivedLicenseConsumptionConfigurationProvisionalConfiguration `json:"provisionalConfigurations,omitempty" yaml:"provisionalConfigurations,omitempty"`

	//
	RenewType string `json:"renewType,omitempty" yaml:"renewType,omitempty"`
}
