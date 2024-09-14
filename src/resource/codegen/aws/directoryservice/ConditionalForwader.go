package directoryservice

type ConditionalForwader struct {
	// The fully qualified domain name of the remote domain for which forwarders will be used.
	RemoteDomainName string `json:"remoteDomainName,omitempty" yaml:"remoteDomainName,omitempty"`

	// ID of directory.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// A list of forwarder IP addresses.
	DnsIps []string `json:"dnsIps,omitempty" yaml:"dnsIps,omitempty"`
}
