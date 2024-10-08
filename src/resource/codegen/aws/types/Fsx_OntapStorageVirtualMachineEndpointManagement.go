package types

type Fsx_OntapStorageVirtualMachineEndpointManagement struct {
	// IP addresses of the storage virtual machine endpoint.
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`

	// The Domain Name Service (DNS) name for the storage virtual machine. You can mount your storage virtual machine using its DNS name.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`
}
