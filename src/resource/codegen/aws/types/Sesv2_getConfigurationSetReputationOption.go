package types

type Sesv2_getConfigurationSetReputationOption struct {
	// The date and time (in Unix time) when the reputation metrics were last given a fresh start.
	LastFreshStart string `json:"lastFreshStart,omitempty" yaml:"lastFreshStart,omitempty"`

	// Specifies whether tracking of reputation metrics is enabled.
	ReputationMetricsEnabled bool `json:"reputationMetricsEnabled,omitempty" yaml:"reputationMetricsEnabled,omitempty"`
}
