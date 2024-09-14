package types

type Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScope struct {
	// Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See Destination below for details.
	Destinations []Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScopeDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// Set of protocols to inspect for, specified using the protocol's assigned internet protocol number (IANA). Network Firewall currently supports TCP only. Valid values: `6`
	Protocols []int `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	// Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See Source Ports below for details.
	SourcePorts []Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScopeSourcePort `json:"sourcePorts,omitempty" yaml:"sourcePorts,omitempty"`

	// Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See Source below for details.
	Sources []Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScopeSource `json:"sources,omitempty" yaml:"sources,omitempty"`

	// Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See Destination Ports below for details.
	DestinationPorts []Networkfirewall_TlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScopeDestinationPort `json:"destinationPorts,omitempty" yaml:"destinationPorts,omitempty"`
}
