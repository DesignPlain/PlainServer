package types

type Fsx_getOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration struct {
	// The fully qualified domain name of the self-managed AD directory.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The name of the domain group whose members have administrative privileges for the FSx file system.
	FileSystemAdministratorsGroup string `json:"fileSystemAdministratorsGroup,omitempty" yaml:"fileSystemAdministratorsGroup,omitempty"`

	// The fully qualified distinguished name of the organizational unit within the self-managed AD directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined.
	OrganizationalUnitDistinguishedName string `json:"organizationalUnitDistinguishedName,omitempty" yaml:"organizationalUnitDistinguishedName,omitempty"`

	// The user name for the service account on your self-managed AD domain that FSx uses to join to your AD domain.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.
	DnsIps []string `json:"dnsIps,omitempty" yaml:"dnsIps,omitempty"`
}
