package types

type Netapp_VolumeExportPolicyRule struct {
	// If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'integrity' kerberos security mode.
	Kerberos5iReadOnly bool `json:"kerberos5iReadOnly,omitempty" yaml:"kerberos5iReadOnly,omitempty"`

	// If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'integrity' kerberos security mode. The 'kerberos5iReadOnly' value is ignored if this is enabled.
	Kerberos5iReadWrite bool `json:"kerberos5iReadWrite,omitempty" yaml:"kerberos5iReadWrite,omitempty"`

	// Enable to apply the export rule to NFSV3 clients.
	Nfsv3 bool `json:"nfsv3,omitempty" yaml:"nfsv3,omitempty"`

	// Enable to apply the export rule to NFSV4.1 clients.
	Nfsv4 bool `json:"nfsv4,omitempty" yaml:"nfsv4,omitempty"`

	/*
	   Defines the access type for clients matching the `allowedClients` specification.
	   Possible values are: `READ_ONLY`, `READ_WRITE`, `READ_NONE`.
	*/
	AccessType string `json:"accessType,omitempty" yaml:"accessType,omitempty"`

	// Defines the client ingress specification (allowed clients) as a comma seperated list with IPv4 CIDRs or IPv4 host addresses.
	AllowedClients string `json:"allowedClients,omitempty" yaml:"allowedClients,omitempty"`

	// If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'authentication' kerberos security mode. The 'kerberos5ReadOnly' value is ignored if this is enabled.
	Kerberos5ReadWrite bool `json:"kerberos5ReadWrite,omitempty" yaml:"kerberos5ReadWrite,omitempty"`

	// If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'privacy' kerberos security mode. The 'kerberos5pReadOnly' value is ignored if this is enabled.
	Kerberos5pReadWrite bool `json:"kerberos5pReadWrite,omitempty" yaml:"kerberos5pReadWrite,omitempty"`

	// If enabled, the root user (UID = 0) of the specified clients doesn't get mapped to nobody (UID = 65534). This is also known as no_root_squash.
	HasRootAccess string `json:"hasRootAccess,omitempty" yaml:"hasRootAccess,omitempty"`

	// If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'authentication' kerberos security mode.
	Kerberos5ReadOnly bool `json:"kerberos5ReadOnly,omitempty" yaml:"kerberos5ReadOnly,omitempty"`

	// If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification. It enables nfs clients to mount using 'privacy' kerberos security mode.
	Kerberos5pReadOnly bool `json:"kerberos5pReadOnly,omitempty" yaml:"kerberos5pReadOnly,omitempty"`
}
