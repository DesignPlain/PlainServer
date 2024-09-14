package ec2

import types "libds/aws/types"

type EipDomainName struct {
	// The allocation ID.
	AllocationId string `json:"allocationId,omitempty" yaml:"allocationId,omitempty"`

	// The domain name to modify for the IP address.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	//
	Timeouts types.Ec2_EipDomainNameTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
