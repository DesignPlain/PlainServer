package types

type Fsx_getOntapFileSystemEndpointManagement struct {
	// DNS name for the file system.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	//
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`
}
