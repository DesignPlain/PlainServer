package types

type Datasync_EfsLocationEc2Config struct {
	// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
	SecurityGroupArns []string `json:"securityGroupArns,omitempty" yaml:"securityGroupArns,omitempty"`

	// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
	SubnetArn string `json:"subnetArn,omitempty" yaml:"subnetArn,omitempty"`
}
