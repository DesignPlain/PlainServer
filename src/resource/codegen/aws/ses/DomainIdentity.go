package ses

type DomainIdentity struct {
	// The domain name to assign to SES
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
