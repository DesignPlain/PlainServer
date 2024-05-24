package types

type Fsx_WindowsFileSystemSelfManagedActiveDirectory struct {
	// A list of up to two IP addresses of DNS servers or domain controllers in the self-managed AD directory. The IP addresses need to be either in the same VPC CIDR range as the file system or in the private IP version 4 (IPv4) address ranges as specified in [RFC 1918](https://tools.ietf.org/html/rfc1918).
	DnsIps []string `json:"dnsIps,omitempty" yaml:"dnsIps,omitempty"`

	// The fully qualified domain name of the self-managed AD directory. For example, `corp.example.com`.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The name of the domain group whose members are granted administrative privileges for the file system. Administrative privileges include taking ownership of files and folders, and setting audit controls (audit ACLs) on files and folders. The group that you specify must already exist in your domain. Defaults to `Domain Admins`.
	FileSystemAdministratorsGroup string `json:"fileSystemAdministratorsGroup,omitempty" yaml:"fileSystemAdministratorsGroup,omitempty"`

	// The fully qualified distinguished name of the organizational unit within your self-managed AD directory that the Windows File Server instance will join. For example, `OU=FSx,DC=yourdomain,DC=corp,DC=com`. Only accepts OU as the direct parent of the file system. If none is provided, the FSx file system is created in the default location of your self-managed AD directory. To learn more, see [RFC 2253](https://tools.ietf.org/html/rfc2253).
	OrganizationalUnitDistinguishedName string `json:"organizationalUnitDistinguishedName,omitempty" yaml:"organizationalUnitDistinguishedName,omitempty"`

	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
