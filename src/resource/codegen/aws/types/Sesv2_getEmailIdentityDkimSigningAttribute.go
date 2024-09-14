package types

type Sesv2_getEmailIdentityDkimSigningAttribute struct {
	//
	DomainSigningSelector string `json:"domainSigningSelector,omitempty" yaml:"domainSigningSelector,omitempty"`

	// [Easy DKIM] The last time a key pair was generated for this identity.
	LastKeyGenerationTimestamp string `json:"lastKeyGenerationTimestamp,omitempty" yaml:"lastKeyGenerationTimestamp,omitempty"`

	// [Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.
	NextSigningKeyLength string `json:"nextSigningKeyLength,omitempty" yaml:"nextSigningKeyLength,omitempty"`

	// A string that indicates how DKIM was configured for the identity. `AWS_SES` indicates that DKIM was configured for the identity by using Easy DKIM. `EXTERNAL` indicates that DKIM was configured for the identity by using Bring Your Own DKIM (BYODKIM).
	SigningAttributesOrigin string `json:"signingAttributesOrigin,omitempty" yaml:"signingAttributesOrigin,omitempty"`

	// Describes whether or not Amazon SES has successfully located the DKIM records in the DNS records for the domain. See the [AWS SES API v2 Reference](https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_DkimAttributes.html#SES-Type-DkimAttributes-Status) for supported statuses.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// If you used Easy DKIM to configure DKIM authentication for the domain, then this object contains a set of unique strings that you use to create a set of CNAME records that you add to the DNS configuration for your domain. When Amazon SES detects these records in the DNS configuration for your domain, the DKIM authentication process is complete. If you configured DKIM authentication for the domain by providing your own public-private key pair, then this object contains the selector for the public key.
	Tokens []string `json:"tokens,omitempty" yaml:"tokens,omitempty"`

	// [Easy DKIM] The key length of the DKIM key pair in use.
	CurrentSigningKeyLength string `json:"currentSigningKeyLength,omitempty" yaml:"currentSigningKeyLength,omitempty"`

	//
	DomainSigningPrivateKey string `json:"domainSigningPrivateKey,omitempty" yaml:"domainSigningPrivateKey,omitempty"`
}
