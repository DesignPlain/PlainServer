package types

type Identityplatform_ConfigMfaProviderConfigTotpProviderConfig struct {
	// The allowed number of adjacent intervals that will be used for verification to avoid clock skew.
	AdjacentIntervals int `json:"adjacentIntervals,omitempty" yaml:"adjacentIntervals,omitempty"`
}
