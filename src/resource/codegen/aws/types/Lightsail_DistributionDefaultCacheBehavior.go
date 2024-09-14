package types

type Lightsail_DistributionDefaultCacheBehavior struct {
	// The cache behavior of the distribution. Valid values: `cache` and `dont-cache`.
	Behavior string `json:"behavior,omitempty" yaml:"behavior,omitempty"`
}
