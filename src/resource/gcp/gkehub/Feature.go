package gkehub

import types "DesignSphere_Server/src/resource/gcp/types"

type Feature struct {
	/*
	   Optional. Fleet Default Membership Configuration.
	   Structure is documented below.
	*/
	FleetDefaultMemberConfig types.Gkehub_FeatureFleetDefaultMemberConfig `json:"fleetDefaultMemberConfig,omitempty" yaml:"fleetDefaultMemberConfig,omitempty"`

	/*
	   GCP labels for this Feature.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location for the resource


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The full, unique name of this Feature resource
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	   Structure is documented below.
	*/
	Spec types.Gkehub_FeatureSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}
