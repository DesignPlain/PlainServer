package types

type Compute_PerInstanceConfigPreservedState struct {
	/*
	   Stateful disks for the instance.
	   Structure is documented below.
	*/
	Disks []Compute_PerInstanceConfigPreservedStateDisk `json:"disks,omitempty" yaml:"disks,omitempty"`

	/*
	   Preserved external IPs defined for this instance. This map is keyed with the name of the network interface.
	   Structure is documented below.
	*/
	ExternalIps []Compute_PerInstanceConfigPreservedStateExternalIp `json:"externalIps,omitempty" yaml:"externalIps,omitempty"`

	/*
	   Preserved internal IPs defined for this instance. This map is keyed with the name of the network interface.
	   Structure is documented below.
	*/
	InternalIps []Compute_PerInstanceConfigPreservedStateInternalIp `json:"internalIps,omitempty" yaml:"internalIps,omitempty"`

	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}
