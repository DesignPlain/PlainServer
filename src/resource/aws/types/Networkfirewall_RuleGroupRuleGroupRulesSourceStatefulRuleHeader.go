package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatefulRuleHeader struct {
	// The protocol to inspect. Valid values: `IP`, `TCP`, `UDP`, `ICMP`, `HTTP`, `FTP`, `TLS`, `SMB`, `DNS`, `DCERPC`, `SSH`, `SMTP`, `IMAP`, `MSN`, `KRB5`, `IKEV2`, `TFTP`, `NTP`, `DHCP`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The source IP address or address range for, in CIDR notation. To match with any address, specify `ANY`.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// The source port to inspect for. To match with any address, specify `ANY`.
	SourcePort string `json:"sourcePort,omitempty" yaml:"sourcePort,omitempty"`

	// The destination IP address or address range to inspect for, in CIDR notation. To match with any address, specify `ANY`.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The destination port to inspect for. To match with any address, specify `ANY`.
	DestinationPort string `json:"destinationPort,omitempty" yaml:"destinationPort,omitempty"`

	// The direction of traffic flow to inspect. Valid values: `ANY` or `FORWARD`.
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`
}
