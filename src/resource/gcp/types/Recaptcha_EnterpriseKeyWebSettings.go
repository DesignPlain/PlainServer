package types

type Recaptcha_EnterpriseKeyWebSettings struct {
	// If set to true, it means allowed_domains will not be enforced.
	AllowAllDomains bool `json:"allowAllDomains,omitempty" yaml:"allowAllDomains,omitempty"`

	// If set to true, the key can be used on AMP (Accelerated Mobile Pages) websites. This is supported only for the SCORE integration type.
	AllowAmpTraffic bool `json:"allowAmpTraffic,omitempty" yaml:"allowAmpTraffic,omitempty"`

	// Domains or subdomains of websites allowed to use the key. All subdomains of an allowed domain are automatically allowed. A valid domain requires a host and must not include any path, port, query or fragment. Examples: 'example.com' or 'subdomain.example.com'
	AllowedDomains []string `json:"allowedDomains,omitempty" yaml:"allowedDomains,omitempty"`

	// Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE. Possible values: CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED, USABILITY, BALANCE, SECURITY
	ChallengeSecurityPreference string `json:"challengeSecurityPreference,omitempty" yaml:"challengeSecurityPreference,omitempty"`

	// Required. Describes how this key is integrated with the website. Possible values: SCORE, CHECKBOX, INVISIBLE
	IntegrationType string `json:"integrationType,omitempty" yaml:"integrationType,omitempty"`
}
