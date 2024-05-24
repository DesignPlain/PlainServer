package licensemanager

type Association struct {
	// ARN of the license configuration.
	LicenseConfigurationArn string `json:"licenseConfigurationArn,omitempty" yaml:"licenseConfigurationArn,omitempty"`

	// ARN of the resource associated with the license configuration.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
