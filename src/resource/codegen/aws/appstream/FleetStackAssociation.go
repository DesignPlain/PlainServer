package appstream

type FleetStackAssociation struct {
	// Name of the fleet.
	FleetName string `json:"fleetName,omitempty" yaml:"fleetName,omitempty"`

	// Name of the stack.
	StackName string `json:"stackName,omitempty" yaml:"stackName,omitempty"`
}
