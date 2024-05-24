package s3outposts

type Endpoint struct {
	// Type of access for the network connectivity. Valid values are `Private` or `CustomerOwnedIp`.
	AccessType string `json:"accessType,omitempty" yaml:"accessType,omitempty"`

	// The ID of a Customer Owned IP Pool. For more on customer owned IP addresses see the [User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/local-rack.html#local-gateway-subnet).
	CustomerOwnedIpv4Pool string `json:"customerOwnedIpv4Pool,omitempty" yaml:"customerOwnedIpv4Pool,omitempty"`

	// Identifier of the Outpost to contain this endpoint.
	OutpostId string `json:"outpostId,omitempty" yaml:"outpostId,omitempty"`

	// Identifier of the EC2 Security Group.
	SecurityGroupId string `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`

	// Identifier of the EC2 Subnet.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
