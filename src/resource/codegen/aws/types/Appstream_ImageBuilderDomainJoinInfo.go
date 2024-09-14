package types

type Appstream_ImageBuilderDomainJoinInfo struct {
	// Fully qualified name of the directory (for example, corp.example.com).
	DirectoryName string `json:"directoryName,omitempty" yaml:"directoryName,omitempty"`

	// Distinguished name of the organizational unit for computer accounts.
	OrganizationalUnitDistinguishedName string `json:"organizationalUnitDistinguishedName,omitempty" yaml:"organizationalUnitDistinguishedName,omitempty"`
}
