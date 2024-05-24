package types

type Cloudfront_DistributionTrustedSigner struct {
	// Whether Origin Shield is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// List of nested attributes for each trusted signer
	Items []Cloudfront_DistributionTrustedSignerItem `json:"items,omitempty" yaml:"items,omitempty"`
}
