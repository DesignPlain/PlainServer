package ec2

type AmiLaunchPermission struct {
	// AWS account ID for the launch permission.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Name of the group for the launch permission. Valid values: `"all"`.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	// ID of the AMI.
	ImageId string `json:"imageId,omitempty" yaml:"imageId,omitempty"`

	// ARN of an organization for the launch permission.
	OrganizationArn string `json:"organizationArn,omitempty" yaml:"organizationArn,omitempty"`

	// ARN of an organizational unit for the launch permission.
	OrganizationalUnitArn string `json:"organizationalUnitArn,omitempty" yaml:"organizationalUnitArn,omitempty"`
}
