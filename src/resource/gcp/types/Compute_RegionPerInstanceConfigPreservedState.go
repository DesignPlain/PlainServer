package types

type Compute_RegionPerInstanceConfigPreservedState struct {
	/*
	   Stateful disks for the instance.
	   Structure is documented below.
	*/
	Disks []Compute_RegionPerInstanceConfigPreservedStateDisk `json:"disks,omitempty" yaml:"disks,omitempty"`

	/*
	   Preserved external IPs defined for this instance. This map is keyed with the name of the network interface.
	   Structure is documented below.
	*/
	ExternalIps []Compute_RegionPerInstanceConfigPreservedStateExternalIp `json:"externalIps,omitempty" yaml:"externalIps,omitempty"`

	/*
	   Preserved internal IPs defined for this instance. This map is keyed with the name of the network interface.
	   Structure is documented below.
	*/
	InternalIps []Compute_RegionPerInstanceConfigPreservedStateInternalIp `json:"internalIps,omitempty" yaml:"internalIps,omitempty"`

	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}
