package apprunner

type CustomDomainAssociation struct {
	// Custom domain endpoint to association. Specify a base domain e.g., `example.com` or a subdomain e.g., `subdomain.example.com`.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Whether to associate the subdomain with the App Runner service in addition to the base domain. Defaults to `true`.
	EnableWwwSubdomain bool `json:"enableWwwSubdomain,omitempty" yaml:"enableWwwSubdomain,omitempty"`

	// ARN of the App Runner service.
	ServiceArn string `json:"serviceArn,omitempty" yaml:"serviceArn,omitempty"`
}
