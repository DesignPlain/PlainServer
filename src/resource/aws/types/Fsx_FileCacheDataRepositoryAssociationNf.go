package types

type Fsx_FileCacheDataRepositoryAssociationNf struct {
	// A list of up to 2 IP addresses of DNS servers used to resolve the NFS file system domain name. The provided IP addresses can either be the IP addresses of a DNS forwarder or resolver that the customer manages and runs inside the customer VPC, or the IP addresses of the on-premises DNS servers.
	DnsIps []string `json:"dnsIps,omitempty" yaml:"dnsIps,omitempty"`

	// The version of the NFS (Network File System) protocol of the NFS data repository. The only supported value is NFS3, which indicates that the data repository must support the NFSv3 protocol. The only supported value is `NFS3`.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
