package types

type Amp_ScraperSource struct {
	// Configuration block for an EKS cluster source. See `eks`.
	Eks Amp_ScraperSourceEks `json:"eks,omitempty" yaml:"eks,omitempty"`
}
