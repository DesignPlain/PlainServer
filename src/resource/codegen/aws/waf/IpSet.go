package waf

import types "libds/aws/types"

type IpSet struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors []types.Waf_IpSetIpSetDescriptor `json:"ipSetDescriptors,omitempty" yaml:"ipSetDescriptors,omitempty"`

	// The name or description of the IPSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
