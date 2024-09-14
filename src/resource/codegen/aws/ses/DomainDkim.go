package ses

type DomainDkim struct {
	// Verified domain name to generate DKIM tokens for.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
