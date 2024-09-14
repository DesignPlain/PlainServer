package types

type Globalaccelerator_ListenerPortRange struct {
	// The first port in the range of ports, inclusive.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// The last port in the range of ports, inclusive.
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`
}
