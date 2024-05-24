package types

type Amp_ScraperDestination struct {
	// Configuration block for an Amazon Managed Prometheus workspace destination. See `amp`.
	Amp Amp_ScraperDestinationAmp `json:"amp,omitempty" yaml:"amp,omitempty"`
}
