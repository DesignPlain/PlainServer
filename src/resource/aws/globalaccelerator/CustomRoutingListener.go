package globalaccelerator

import types "DesignSphere_Server/src/resource/aws/types"

type CustomRoutingListener struct {
	// The Amazon Resource Name (ARN) of a custom routing accelerator.
	AcceleratorArn string `json:"acceleratorArn,omitempty" yaml:"acceleratorArn,omitempty"`

	// The list of port ranges for the connections from clients to the accelerator. Fields documented below.
	PortRanges []types.Globalaccelerator_CustomRoutingListenerPortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`
}
