package types

type Cloudfront_DistributionOriginOriginShield struct {
	// `true` if any of the AWS accounts listed as trusted signers have active CloudFront key pairs
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// AWS Region for Origin Shield. To specify a region, use the region code, not the region name. For example, specify the US East (Ohio) region as `us-east-2`.
	OriginShieldRegion string `json:"originShieldRegion,omitempty" yaml:"originShieldRegion,omitempty"`
}
