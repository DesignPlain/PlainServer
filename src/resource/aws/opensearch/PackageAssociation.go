package opensearch

type PackageAssociation struct {
	// Name of the domain to associate the package with.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Internal ID of the package to associate with a domain.
	PackageId string `json:"packageId,omitempty" yaml:"packageId,omitempty"`
}
