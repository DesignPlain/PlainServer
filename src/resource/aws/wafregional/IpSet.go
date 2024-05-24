package wafregional

import types "DesignSphere_Server/src/resource/aws/types"

type IpSet struct {
	// The name or description of the IPSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors []types.Wafregional_IpSetIpSetDescriptor `json:"ipSetDescriptors,omitempty" yaml:"ipSetDescriptors,omitempty"`
}
