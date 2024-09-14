package types

type Amplify_DomainAssociationSubDomain struct {
	// Branch name setting for the subdomain.
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// DNS record for the subdomain in a space-prefixed and space-delimited format (` CNAME <target>`).
	DnsRecord string `json:"dnsRecord,omitempty" yaml:"dnsRecord,omitempty"`

	// Prefix setting for the subdomain.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Verified status of the subdomain.
	Verified bool `json:"verified,omitempty" yaml:"verified,omitempty"`
}
