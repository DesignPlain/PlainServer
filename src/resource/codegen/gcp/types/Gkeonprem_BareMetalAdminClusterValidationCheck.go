package types

type Gkeonprem_BareMetalAdminClusterValidationCheck struct {
	/*
	   (Output)
	   Options used for the validation check.
	*/
	Options string `json:"options,omitempty" yaml:"options,omitempty"`

	/*
	   (Output)
	   The scenario when the preflight checks were run..
	*/
	Scenario string `json:"scenario,omitempty" yaml:"scenario,omitempty"`

	/*
	   (Output)
	   Specifies the detailed validation check status
	   Structure is documented below.
	*/
	Statuses []Gkeonprem_BareMetalAdminClusterValidationCheckStatus `json:"statuses,omitempty" yaml:"statuses,omitempty"`
}
