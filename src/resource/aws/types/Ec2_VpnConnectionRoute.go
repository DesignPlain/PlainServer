package types

type Ec2_VpnConnectionRoute struct {
	// The CIDR block associated with the local subnet of the customer data center.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`

	// Indicates how the routes were provided.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// The current state of the static route.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
