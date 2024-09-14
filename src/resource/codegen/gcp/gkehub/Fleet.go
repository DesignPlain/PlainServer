package gkehub

import types "libds/gcp/types"

type Fleet struct {
	/*
	   The default cluster configurations to apply across the fleet.
	   Structure is documented below.
	*/
	DefaultClusterConfig types.Gkehub_FleetDefaultClusterConfig `json:"defaultClusterConfig,omitempty" yaml:"defaultClusterConfig,omitempty"`

	/*
	   A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
	   Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
