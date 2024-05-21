package types

type Osconfig_PatchDeploymentInstanceFilterGroupLabel struct {
	/*
	   Compute Engine instance labels that must be present for a VM instance to be targeted by this filter

	   - - -
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
