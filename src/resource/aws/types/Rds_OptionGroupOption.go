package types

type Rds_OptionGroupOption struct {
	// Version of the option (e.g., 13.1.0.0). Leaving out or removing `version` from your configuration does not remove or clear a version from the option in AWS. AWS may assign a default version. Not including `version` in your configuration means that the AWS provider will ignore a previously set value, a value set by AWS, and any version changes.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// List of VPC Security Groups for which the option is enabled.
	VpcSecurityGroupMemberships []string `json:"vpcSecurityGroupMemberships,omitempty" yaml:"vpcSecurityGroupMemberships,omitempty"`

	// List of DB Security Groups for which the option is enabled.
	DbSecurityGroupMemberships []string `json:"dbSecurityGroupMemberships,omitempty" yaml:"dbSecurityGroupMemberships,omitempty"`

	// Name of the option (e.g., MEMCACHED).
	OptionName string `json:"optionName,omitempty" yaml:"optionName,omitempty"`

	// The option settings to apply. See `option_settings` Block below for more details.
	OptionSettings []Rds_OptionGroupOptionOptionSetting `json:"optionSettings,omitempty" yaml:"optionSettings,omitempty"`

	// Port number when connecting to the option (e.g., 11211). Leaving out or removing `port` from your configuration does not remove or clear a port from the option in AWS. AWS may assign a default port. Not including `port` in your configuration means that the AWS provider will ignore a previously set value, a value set by AWS, and any port changes.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
