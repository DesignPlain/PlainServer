package licensemanager

type LicenseConfiguration struct {
	// Name of the license configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of the license configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Number of licenses managed by the license configuration.
	LicenseCount int `json:"licenseCount,omitempty" yaml:"licenseCount,omitempty"`

	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit bool `json:"licenseCountHardLimit,omitempty" yaml:"licenseCountHardLimit,omitempty"`

	// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
	LicenseCountingType string `json:"licenseCountingType,omitempty" yaml:"licenseCountingType,omitempty"`

	// Array of configured License Manager rules.
	LicenseRules []string `json:"licenseRules,omitempty" yaml:"licenseRules,omitempty"`
}
