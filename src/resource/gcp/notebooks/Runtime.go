package notebooks

import types "DesignSphere_Server/src/resource/gcp/types"

type Runtime struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The config settings for software inside the runtime.
	   Structure is documented below.
	*/
	SoftwareConfig types.Notebooks_RuntimeSoftwareConfig `json:"softwareConfig,omitempty" yaml:"softwareConfig,omitempty"`

	/*
	   Use a Compute Engine VM image to start the managed notebook instance.
	   Structure is documented below.
	*/
	VirtualMachine types.Notebooks_RuntimeVirtualMachine `json:"virtualMachine,omitempty" yaml:"virtualMachine,omitempty"`

	/*
	   The config settings for accessing runtime.
	   Structure is documented below.
	*/
	AccessConfig types.Notebooks_RuntimeAccessConfig `json:"accessConfig,omitempty" yaml:"accessConfig,omitempty"`

	/*
	   The labels to associate with this runtime. Label --keys-- must
	   contain 1 to 63 characters, and must conform to [RFC 1035]
	   (https://www.ietf.org/rfc/rfc1035.txt). Label --values-- may be
	   empty, but, if present, must contain 1 to 63 characters, and must
	   conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
	   more than 32 labels can be associated with a cluster.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   A reference to the zone where the machine resides.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name specified for the Notebook runtime.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
