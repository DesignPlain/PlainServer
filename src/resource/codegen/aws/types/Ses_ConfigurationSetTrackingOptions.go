package types

type Ses_ConfigurationSetTrackingOptions struct {
	// Custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
	CustomRedirectDomain string `json:"customRedirectDomain,omitempty" yaml:"customRedirectDomain,omitempty"`
}
