package types

type Licensemanager_getReceivedLicenseConsumptionConfigurationBorrowConfiguration struct {
	// Indicates whether early check-ins are allowed.
	AllowEarlyCheckIn bool `json:"allowEarlyCheckIn,omitempty" yaml:"allowEarlyCheckIn,omitempty"`

	// Maximum time for the provisional configuration, in minutes.
	MaxTimeToLiveInMinutes int `json:"maxTimeToLiveInMinutes,omitempty" yaml:"maxTimeToLiveInMinutes,omitempty"`
}
