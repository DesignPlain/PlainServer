package types

type Fsx_getOntapFileSystemEndpointIntercluster struct {
	// DNS name for the file system (e.g. `fs-12345678.corp.example.com`).
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	//
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`
}
