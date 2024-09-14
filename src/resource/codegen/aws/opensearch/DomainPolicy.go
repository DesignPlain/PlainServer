package opensearch

type DomainPolicy struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies string `json:"accessPolicies,omitempty" yaml:"accessPolicies,omitempty"`

	// Name of the domain.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
