package wafregional

import types "libds/aws/types"

type IpSet struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors []types.Wafregional_IpSetIpSetDescriptor `json:"ipSetDescriptors,omitempty" yaml:"ipSetDescriptors,omitempty"`

	// The name or description of the IPSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
