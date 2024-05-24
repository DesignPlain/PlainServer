package appsync

type DomainNameApiAssociation struct {
	// API ID.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// Appsync domain name.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
