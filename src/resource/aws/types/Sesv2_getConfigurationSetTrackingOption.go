package types

type Sesv2_getConfigurationSetTrackingOption struct {
	// The domain to use for tracking open and click events.
	CustomRedirectDomain string `json:"customRedirectDomain,omitempty" yaml:"customRedirectDomain,omitempty"`
}
