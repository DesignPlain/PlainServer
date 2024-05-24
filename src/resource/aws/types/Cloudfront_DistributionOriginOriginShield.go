package types

type Cloudfront_DistributionOriginOriginShield struct {
	// Whether Origin Shield is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// AWS Region for Origin Shield. To specify a region, use the region code, not the region name. For example, specify the US East (Ohio) region as `us-east-2`.
	OriginShieldRegion string `json:"originShieldRegion,omitempty" yaml:"originShieldRegion,omitempty"`
}
