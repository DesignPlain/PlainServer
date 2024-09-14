package ec2

type NetworkInterfaceSecurityGroupAttachment struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The ID of the security group.
	SecurityGroupId string `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`
}
