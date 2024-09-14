package types

type Cloudfront_DistributionTrustedKeyGroup struct {
	// `true` if any of the AWS accounts listed as trusted signers have active CloudFront key pairs
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// List of nested attributes for each trusted signer
	Items []Cloudfront_DistributionTrustedKeyGroupItem `json:"items,omitempty" yaml:"items,omitempty"`
}
