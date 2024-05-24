package types

type Lightsail_DistributionCacheBehaviorSettingsForwardedQueryStrings struct {
	// Indicates whether the distribution forwards and caches based on query strings.
	Option bool `json:"option,omitempty" yaml:"option,omitempty"`

	// The specific query strings that the distribution forwards to the origin.
	QueryStringsAllowedLists []string `json:"queryStringsAllowedLists,omitempty" yaml:"queryStringsAllowedLists,omitempty"`
}
