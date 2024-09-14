package types

type Gkeonprem_BareMetalAdminClusterValidationCheckStatus struct {
	/*
	   (Output)
	   Individual checks which failed as part of the Preflight check execution.
	   Structure is documented below.
	*/
	Results []Gkeonprem_BareMetalAdminClusterValidationCheckStatusResult `json:"results,omitempty" yaml:"results,omitempty"`
}
