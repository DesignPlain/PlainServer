package types

type Licensemanager_getReceivedLicenseEntitlement struct {
	// The key name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Entitlement unit.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// The value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Indicates whether check-ins are allowed.
	AllowCheckIn bool `json:"allowCheckIn,omitempty" yaml:"allowCheckIn,omitempty"`

	// Maximum entitlement count. Use if the unit is not None.
	MaxCount int `json:"maxCount,omitempty" yaml:"maxCount,omitempty"`
}
