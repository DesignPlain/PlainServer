package types

type Opensearch_DomainOffPeakWindowOptions struct {
	// Enabled disabled toggle for off-peak update window.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	OffPeakWindow Opensearch_DomainOffPeakWindowOptionsOffPeakWindow `json:"offPeakWindow,omitempty" yaml:"offPeakWindow,omitempty"`
}
