package types

type Compute_PacketMirroringMirroredResources struct {
	/*
	   All the listed instances will be mirrored.  Specify at most 50.
	   Structure is documented below.
	*/
	Instances []Compute_PacketMirroringMirroredResourcesInstance `json:"instances,omitempty" yaml:"instances,omitempty"`

	/*
	   All instances in one of these subnetworks will be mirrored.
	   Structure is documented below.
	*/
	Subnetworks []Compute_PacketMirroringMirroredResourcesSubnetwork `json:"subnetworks,omitempty" yaml:"subnetworks,omitempty"`

	// All instances with these tags will be mirrored.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
