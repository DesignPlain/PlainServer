package types

type Gkeonprem_VMwareClusterNetworkConfigHostConfig struct {
	/*
	   DNS search domains.

	   <a name="nested_control_plane_v2_config"></a>The `control_plane_v2_config` block supports:
	*/
	DnsSearchDomains []string `json:"dnsSearchDomains,omitempty" yaml:"dnsSearchDomains,omitempty"`

	// DNS servers.
	DnsServers []string `json:"dnsServers,omitempty" yaml:"dnsServers,omitempty"`

	// NTP servers.
	NtpServers []string `json:"ntpServers,omitempty" yaml:"ntpServers,omitempty"`
}
