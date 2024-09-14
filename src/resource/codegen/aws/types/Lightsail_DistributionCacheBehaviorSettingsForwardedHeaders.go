package types

type Lightsail_DistributionCacheBehaviorSettingsForwardedHeaders struct {
	// The specific headers to forward to your distribution's origin.
	HeadersAllowLists []string `json:"headersAllowLists,omitempty" yaml:"headersAllowLists,omitempty"`

	// The headers that you want your distribution to forward to your origin and base caching on.
	Option string `json:"option,omitempty" yaml:"option,omitempty"`
}
