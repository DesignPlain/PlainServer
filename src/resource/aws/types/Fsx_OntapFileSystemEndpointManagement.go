package types

type Fsx_OntapFileSystemEndpointManagement struct {
	// The Domain Name Service (DNS) name for the file system. You can mount your file system using its DNS name.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	// IP addresses of the file system endpoint.
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`
}
