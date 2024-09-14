package types

type Gkeonprem_BareMetalClusterValidationCheckStatus struct {
	/*
	   (Output)
	   Individual checks which failed as part of the Preflight check execution.
	   Structure is documented below.
	*/
	Results []Gkeonprem_BareMetalClusterValidationCheckStatusResult `json:"results,omitempty" yaml:"results,omitempty"`
}
