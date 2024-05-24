package types

type Gamelift_FleetEc2InboundPermission struct {
	// Starting value for a range of allowed port numbers.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// Range of allowed IP addresses expressed in CIDR notationE.g., `000.000.000.000/[subnet mask]` or `0.0.0.0/[subnet mask]`.
	IpRange string `json:"ipRange,omitempty" yaml:"ipRange,omitempty"`

	// Network communication protocol used by the fleetE.g., `TCP` or `UDP`
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than `from_port`.
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`
}
