package types

type Opensearch_getDomainOffPeakWindowOptions struct {
	// Enabled disabled toggle for off-peak update window
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	OffPeakWindows []Opensearch_getDomainOffPeakWindowOptionsOffPeakWindow `json:"offPeakWindows,omitempty" yaml:"offPeakWindows,omitempty"`
}
