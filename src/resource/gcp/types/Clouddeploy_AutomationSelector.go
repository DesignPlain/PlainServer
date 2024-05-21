package types

type Clouddeploy_AutomationSelector struct {
	/*
	   Contains attributes about a target.
	   Structure is documented below.
	*/
	Targets []Clouddeploy_AutomationSelectorTarget `json:"targets,omitempty" yaml:"targets,omitempty"`
}
