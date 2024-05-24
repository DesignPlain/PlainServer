package types

type Opensearch_DomainOffPeakWindowOptionsOffPeakWindowWindowStartTime struct {
	// Starting minute of the 10-hour window for updates
	Minutes int `json:"minutes,omitempty" yaml:"minutes,omitempty"`

	// Starting hour of the 10-hour window for updates
	Hours int `json:"hours,omitempty" yaml:"hours,omitempty"`
}
