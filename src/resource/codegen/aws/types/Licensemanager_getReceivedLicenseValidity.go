package types

type Licensemanager_getReceivedLicenseValidity struct {
	// Start of the validity time range.
	Begin string `json:"begin,omitempty" yaml:"begin,omitempty"`

	// End of the validity time range.
	End string `json:"end,omitempty" yaml:"end,omitempty"`
}
