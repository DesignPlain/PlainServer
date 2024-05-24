package types

type Imagebuilder_DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermission struct {
	// Set of AWS Organization ARNs to assign.
	OrganizationArns []string `json:"organizationArns,omitempty" yaml:"organizationArns,omitempty"`

	// Set of AWS Organizational Unit ARNs to assign.
	OrganizationalUnitArns []string `json:"organizationalUnitArns,omitempty" yaml:"organizationalUnitArns,omitempty"`

	// Set of EC2 launch permission user groups to assign. Use `all` to distribute a public AMI.
	UserGroups []string `json:"userGroups,omitempty" yaml:"userGroups,omitempty"`

	// Set of AWS Account identifiers to assign.
	UserIds []string `json:"userIds,omitempty" yaml:"userIds,omitempty"`
}
