package types

type Gkeonprem_VMwareClusterValidationCheckStatus struct {
	/*
	   (Output)
	   Individual checks which failed as part of the Preflight check execution.
	   Structure is documented below.
	*/
	Results []Gkeonprem_VMwareClusterValidationCheckStatusResult `json:"results,omitempty" yaml:"results,omitempty"`
}
