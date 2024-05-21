package vmwareengine

type Network struct {
	// User-provided description for this VMware Engine network.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The location where the VMwareEngineNetwork should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the VMwareEngineNetwork.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   VMware Engine network type.
	   Possible values are: `LEGACY`, `STANDARD`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
