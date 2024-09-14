package types

type Networkconnectivity_SpokeLinkedRouterApplianceInstancesInstance struct {
	// The IP address on the VM to use for peering.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   The URI of the virtual machine resource

	   - - -
	*/
	VirtualMachine string `json:"virtualMachine,omitempty" yaml:"virtualMachine,omitempty"`
}
