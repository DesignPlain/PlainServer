package ses

type DomainIdentityVerification struct {
	// The domain name of the SES domain identity to verify.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
