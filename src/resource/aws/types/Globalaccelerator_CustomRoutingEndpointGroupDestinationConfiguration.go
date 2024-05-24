package types

type Globalaccelerator_CustomRoutingEndpointGroupDestinationConfiguration struct {
	// The protocol for the endpoint group that is associated with a custom routing accelerator. The protocol can be either `"TCP"` or `"UDP"`.
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	// The last port, inclusive, in the range of ports for the endpoint group that is associated with a custom routing accelerator.
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// The first port, inclusive, in the range of ports for the endpoint group that is associated with a custom routing accelerator.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`
}
