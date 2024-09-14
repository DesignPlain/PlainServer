package types

type Securityhub_OrganizationConfigurationOrganizationConfiguration struct {
	// Indicates whether the organization uses local or central configuration. If using central configuration, `auto_enable` must be set to `false` and `auto_enable_standards` set to `NONE`. More information can be found in the [documentation for central configuration](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html). Valid values: `LOCAL`, `CENTRAL`.
	ConfigurationType string `json:"configurationType,omitempty" yaml:"configurationType,omitempty"`
}
