package types

type Finspace_KxEnvironmentCustomDnsConfiguration struct {
	// IP address of the DNS server.
	CustomDnsServerIp string `json:"customDnsServerIp,omitempty" yaml:"customDnsServerIp,omitempty"`

	// Name of the DNS server.
	CustomDnsServerName string `json:"customDnsServerName,omitempty" yaml:"customDnsServerName,omitempty"`
}
