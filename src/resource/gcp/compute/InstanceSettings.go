package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type InstanceSettings struct {
	/*
	   The metadata key/value pairs assigned to all the instances in the corresponding scope.
	   Structure is documented below.
	*/
	Metadata types.Compute_InstanceSettingsMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A reference to the zone where the machine resides.


	   - - -
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
