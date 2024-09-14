package ec2

import types "libds/aws/types"

type VpcIpamPoolCidr struct {
	// If provided, the cidr provisioned into the specified pool will be the next available cidr given this declared netmask length. Conflicts with `cidr`.
	NetmaskLength int `json:"netmaskLength,omitempty" yaml:"netmaskLength,omitempty"`

	// The CIDR you want to assign to the pool. Conflicts with `netmask_length`.
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	// A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP. This is not stored in the state file. See cidr_authorization_context for more information.
	CidrAuthorizationContext types.Ec2_VpcIpamPoolCidrCidrAuthorizationContext `json:"cidrAuthorizationContext,omitempty" yaml:"cidrAuthorizationContext,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId string `json:"ipamPoolId,omitempty" yaml:"ipamPoolId,omitempty"`
}
