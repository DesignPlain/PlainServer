package types

type Imagebuilder_getDistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermission struct {
	// Set of AWS Organization ARNs.
	OrganizationArns []string `json:"organizationArns,omitempty" yaml:"organizationArns,omitempty"`

	// Set of AWS Organizational Unit ARNs.
	OrganizationalUnitArns []string `json:"organizationalUnitArns,omitempty" yaml:"organizationalUnitArns,omitempty"`

	// Set of EC2 launch permission user groups.
	UserGroups []string `json:"userGroups,omitempty" yaml:"userGroups,omitempty"`

	// Set of AWS Account identifiers.
	UserIds []string `json:"userIds,omitempty" yaml:"userIds,omitempty"`
}
