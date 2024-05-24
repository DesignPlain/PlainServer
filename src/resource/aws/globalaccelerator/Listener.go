package globalaccelerator

import types "DesignSphere_Server/src/resource/aws/types"

type Listener struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	AcceleratorArn string `json:"acceleratorArn,omitempty" yaml:"acceleratorArn,omitempty"`

	// Direct all requests from a user to the same endpoint. Valid values are `NONE`, `SOURCE_IP`. Default: `NONE`. If `NONE`, Global Accelerator uses the "five-tuple" properties of source IP address, source port, destination IP address, destination port, and protocol to select the hash value. If `SOURCE_IP`, Global Accelerator uses the "two-tuple" properties of source (client) IP address and destination IP address to select the hash value.
	ClientAffinity string `json:"clientAffinity,omitempty" yaml:"clientAffinity,omitempty"`

	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges []types.Globalaccelerator_ListenerPortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`

	// The protocol for the connections from clients to the accelerator. Valid values are `TCP`, `UDP`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
