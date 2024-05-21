package types

type Monitoring_SloBasicSliAvailability struct {
	// Whether an availability SLI is enabled or not. Must be set to `true. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
