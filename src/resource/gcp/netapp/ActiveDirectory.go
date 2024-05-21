package netapp

type ActiveDirectory struct {
	// Fully qualified domain name for the Active Directory domain.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// Hostname of the Active Directory server used as Kerberos Key Distribution Center. Only requried for volumes using kerberized NFSv4.1
	KdcHostname string `json:"kdcHostname,omitempty" yaml:"kdcHostname,omitempty"`

	// IP address of the Active Directory server used as Kerberos Key Distribution Center.
	KdcIp string `json:"kdcIp,omitempty" yaml:"kdcIp,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be signed.
	LdapSigning bool `json:"ldapSigning,omitempty" yaml:"ldapSigning,omitempty"`

	// Name of the region for the policy to apply to.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Password for specified username. Note - Manual changes done to the password will not be detected. Terraform will not
	   re-apply the password, unless you use a new password in Terraform.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	/*
	   NetBIOS name prefix of the server to be created.
	   A five-character random ID is generated automatically, for example, -6f9a, and appended to the prefix. The full UNC share path will have the following format:
	   `\\NetBIOS_PREFIX-ABCD.DOMAIN_NAME\SHARE_NAME`
	*/
	NetBiosPrefix string `json:"netBiosPrefix,omitempty" yaml:"netBiosPrefix,omitempty"`

	/*
	   Local UNIX users on clients without valid user information in Active Directory are blocked from access to LDAP enabled volumes.
	   This option can be used to temporarily switch such volumes to AUTH_SYS authentication (user ID + 1-16 groups).
	*/
	NfsUsersWithLdap bool `json:"nfsUsersWithLdap,omitempty" yaml:"nfsUsersWithLdap,omitempty"`

	// Domain accounts that require elevated privileges such as `SeSecurityPrivilege` to manage security logs. Comma-separated list.
	SecurityOperators []string `json:"securityOperators,omitempty" yaml:"securityOperators,omitempty"`

	// Username for the Active Directory account with permissions to create the compute account within the specified organizational unit.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// Enables AES-128 and AES-256 encryption for Kerberos-based communication with Active Directory.
	AesEncryption bool `json:"aesEncryption,omitempty" yaml:"aesEncryption,omitempty"`

	// Domain user/group accounts to be added to the Backup Operators group of the SMB service. The Backup Operators group allows members to backup and restore files regardless of whether they have read or write access to the files. Comma-separated list.
	BackupOperators []string `json:"backupOperators,omitempty" yaml:"backupOperators,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Comma separated list of DNS server IP addresses for the Active Directory domain.
	Dns string `json:"dns,omitempty" yaml:"dns,omitempty"`

	/*
	   Labels as key value pairs. Example: `{ "owner": "Bob", "department": "finance", "purpose": "testing" }`.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The resource name of the Active Directory pool. Needs to be unique per location.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Name of the Organizational Unit where you intend to create the computer account for NetApp Volumes.
	   Defaults to `CN=Computers` if left empty.
	*/
	OrganizationalUnit string `json:"organizationalUnit,omitempty" yaml:"organizationalUnit,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies an Active Directory site to manage domain controller selection.
	   Use when Active Directory domain controllers in multiple regions are configured. Defaults to `Default-First-Site-Name` if left empty.
	*/
	Site string `json:"site,omitempty" yaml:"site,omitempty"`

	// If enabled, traffic between the SMB server to Domain Controller (DC) will be encrypted.
	EncryptDcConnections bool `json:"encryptDcConnections,omitempty" yaml:"encryptDcConnections,omitempty"`
}
